package domain

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Declare your app error here
// todo: create custom error response
var (
	ErrUserNotFound = echo.NewHTTPError(http.StatusNotFound, "User not found")
	// ...
	// ...
)
