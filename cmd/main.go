package main

import (
	"echo-hello/bootstrap"
	"echo-hello/delivery/route"
	"echo-hello/repository"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Db

	e := echo.New()

	route.Setup(env, db, e)

	res, _ := repository.NewUserRepository(db).FindAll()
	fmt.Println(res)

	e.Logger.Fatal(e.Start(":" + env.Config.ServerPort))
}
