package validation

import (
	"net/http"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomValidator(t *testing.T) {
	v := validator.New()
	cv := NewCustomValidator(v)
	assert.NotNil(t, cv)
	assert.Equal(t, v, cv.validator)
}

func TestCustomValidator_Validate(t *testing.T) {
	// Create a struct for testing
	type TestStruct struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"email"`
	}

	// Create a custom validator with a validator instance
	v := validator.New()
	cv := NewCustomValidator(v)

	t.Run("ValidInput", func(t *testing.T) {
		// Valid input should not return an error
		input := TestStruct{Name: "John", Email: "john@example.com"}
		err := cv.Validate(input)
		assert.NoError(t, err)
	})

	t.Run("InvalidInput", func(t *testing.T) {
		// Invalid input should return an echo.HTTPError with status code 400
		input := TestStruct{Name: "", Email: "invalid-email"}
		err := cv.Validate(input)
		assert.Error(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, http.StatusBadRequest, echoErr.Code)
	})
}

func TestCustomValidator_Validate_TagNameFunc(t *testing.T) {
	// Create a struct for testing
	type TestStruct struct {
		Name string `json:"-"`
	}

	// Create a custom validator with a validator instance
	v := validator.New()
	cv := NewCustomValidator(v)

	t.Run("TagIgnoreField", func(t *testing.T) {
		// The field with json:"-" should be ignored by the validator
		input := TestStruct{Name: "John"}
		err := cv.Validate(input)
		assert.NoError(t, err)
	})

	t.Run("TagIgnoreField_Error", func(t *testing.T) {
		// The field with json:"-" should be ignored by the validator
		input := TestStruct{Name: ""}
		err := cv.Validate(input)
		assert.NoError(t, err)
	})
}
