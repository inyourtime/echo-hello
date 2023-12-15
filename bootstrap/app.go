package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *Env
	Db  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewPgDatabase(app.Env)
	return *app
}
