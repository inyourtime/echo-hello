package route

import (
	"echo-hello/bootstrap"
	"echo-hello/internal/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, e *echo.Echo) {
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

	if env.Config.Env == "development" {
		e.Use(middleware.Logger())
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	logger.Info("[ROUTE] All routes has been register")
}
