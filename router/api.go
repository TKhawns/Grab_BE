package router

import (
	"go_grab/handler"
	middleware "go_grab/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo         *echo.Echo
	UserHandler  handler.UserHandler
	AdminHandler handler.AdminHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/login", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	api.Echo.POST("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
	api.Echo.POST("/user/upload-model", api.UserHandler.HandleUpload)
	api.Echo.POST("/user/get-model", api.UserHandler.GetModelById)
	api.Echo.GET("/admin/get-usermodel", api.AdminHandler.GetUserModel)
	api.Echo.POST("/admin/update-status", api.AdminHandler.UpdateStatus)

}
