package main

import (
	"echo-hello/bootstrap"
	"echo-hello/delivery/route"

	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Db

	e := echo.New()

	route.New(env, db, e)

	e.Logger.Fatal(e.Start(":" + env.Config.ServerPort))
}
