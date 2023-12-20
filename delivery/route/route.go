package route

import (
	"echo-hello/bootstrap"
	"echo-hello/delivery/validation"
	"echo-hello/internal/logger"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type UserReq struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func New(env *bootstrap.Env, db *gorm.DB, e *echo.Echo) {
	/* Middleware */
	e.Use(middleware.CORS())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			logger.Error(err)
			return c.JSON(500, map[string]interface{}{
				"status":  500,
				"message": "Internal Server Error",
			})
		},
	}))

	// Add validator
	e.Validator = validation.NewCustomValidator(validator.New())

	if env.Config.Env == "development" {
		e.Use(middleware.Logger())
	}

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello, World!</h1>")
	})

	e.POST("/", func(c echo.Context) error {
		user := new(UserReq)
		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(user); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	})

	logger.Info("[ROUTE] All routes has been register")
}
