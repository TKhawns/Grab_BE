package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HanldeSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":    "Khanh",
		"address": "VNU",
	})
}
