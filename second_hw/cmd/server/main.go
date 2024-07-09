package main

import (
	"awesomeProject/accounts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.DELETE("/account", accountsHandler.DeleteAccount)
	e.PATCH("/account", accountsHandler.PathAccount)
	e.PUT("/account", accountsHandler.ChangeAccount)

	e.Logger.Fatal(e.Start(":1323"))
}
