package authenticate

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestExtractTokenFromHeader(t *testing.T) {
	cases := []struct {
		name          string
		headerValue   string
		expectedToken string
		expectedError error
	}{
		{
			name:          "Missing Authorization header",
			headerValue:   "",
			expectedToken: "",
			expectedError: errHeaderExtractorValueMissing,
		},
		{
			name:          "Invalid Authorization header format",
			headerValue:   "invalid_format",
			expectedToken: "",
			expectedError: errHeaderExtractorValueInvalid,
		},
		{
			name:          "Valid Authorization header",
			headerValue:   "Bearer valid_token",
			expectedToken: "valid_token",
			expectedError: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			if tc.headerValue != "" {
				req.Header.Set("Authorization", tc.headerValue)
			}

			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			token, err := extractTokenFromHeader(c)

			assert.Equal(t, tc.expectedToken, token)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestJWTMiddleware(t *testing.T) {
	cases := []struct {
		name             string
		authorization    string
		expectedStatus   int
		expectedResponse map[string]interface{}
	}{
		{
			name:             "Missing Token in Authorization Header",
			authorization:    "",
			expectedStatus:   http.StatusUnauthorized,
			expectedResponse: map[string]interface{}{"message": errHeaderExtractorValueMissing.Error()},
		},
		{
			name:             "Invalid Token in Authorization Header",
			authorization:    "invalid_token",
			expectedStatus:   http.StatusUnauthorized,
			expectedResponse: map[string]interface{}{"message": errHeaderExtractorValueInvalid.Error()},
		},
		{
			name:             "Valid Token in Authorization Header",
			authorization:    "Bearer valid_token",
			expectedStatus:   http.StatusOK,
			expectedResponse: map[string]interface{}{"message": "success"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()

			e.GET("/protected", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"message": "success",
				})
			}, JWT())

			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			if tc.authorization != "" {
				req.Header.Set("Authorization", tc.authorization)
			}
			res := httptest.NewRecorder()

			e.ServeHTTP(res, req)

			assert.Equal(t, tc.expectedStatus, res.Code)
			var responseBody map[string]interface{}
			err := json.Unmarshal(res.Body.Bytes(), &responseBody)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedResponse, responseBody)
		})
	}
}
