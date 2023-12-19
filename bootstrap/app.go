package bootstrap

import (
	"echo-hello/internal/logger"

	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Db  *gorm.DB
}

func App() Application {
	logger.New()
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewPgDatabase(app.Env)
	return *app
}
