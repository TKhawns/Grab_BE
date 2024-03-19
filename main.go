package main

import (
	"go_grab/db"
	"go_grab/handler"
	"go_grab/helper"
	"go_grab/repository/repo_implement"
	"go_grab/router"

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

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	UserHandler := handler.UserHandler{
		UserRepo: repo_implement.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: UserHandler,
	}
	api.SetupRouter()

	e.Logger.Fatal(e.Start(":1323"))
}
