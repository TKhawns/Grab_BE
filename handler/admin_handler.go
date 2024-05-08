package handler

import (
	"go_grab/model"
	"go_grab/model/request"

	"go_grab/repository"

	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminRepo repository.AdminRepo
}

func (u *AdminHandler) GetUserModel(c echo.Context) error {

	list_model, err := u.AdminRepo.GetUserModel(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       list_model,
	})
}

func (u *AdminHandler) UpdateStatus(c echo.Context) error {
	req := request.RequestUpdateStatus{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	err := u.AdminRepo.UpdateStatus(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       "OK",
	})
}
