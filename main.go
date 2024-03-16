package main

import (
	"go_grab/db"
	"go_grab/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "postgres",
		Password: "10012003",
		DbName:   "grab_database",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.GET("/", welcome)
	e.GET("/user/sign-in", handler.HanldeSignIn)
	e.Logger.Fatal(e.Start(":1323"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my app")
}
