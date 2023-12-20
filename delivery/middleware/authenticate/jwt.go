package authenticate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

var errHeaderExtractorValueMissing = errors.New("missing value in request header")
var errHeaderExtractorValueInvalid = errors.New("invalid value in request header")

func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t, err := extractTokenFromHeader(c)
			if err != nil {
				return echo.NewHTTPError(401, err.Error())
			}
			fmt.Println(t)
			return next(c)
		}
	}
}

func extractTokenFromHeader(c echo.Context) (string, error) {
	values := c.Request().Header.Values("Authorization")
	if len(values) == 0 {
		return "", errHeaderExtractorValueMissing
	}
	auth := c.Request().Header.Values("Authorization")[0]
	parts := strings.Fields(auth)
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1], nil
	}
	return "", errHeaderExtractorValueInvalid
}
