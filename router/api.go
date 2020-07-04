package router

import (
	"github-trending/handler"
	middleware "github-trending/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {

	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)


	user := api.Echo.Group("/user",middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/profile/update",api.UserHandler.UpdateProfile)


	github := api.Echo.Group("/github",middleware.JWTMiddleware())
	github.GET("/trending",api.RepoHandler.RepoTrending)

	bookmark := api.Echo.Group("bookmark",middleware.JWTMiddleware())
	bookmark.GET("/list",api.RepoHandler.SelectBookmarks)
	bookmark.POST("/add",api.RepoHandler.Bookmark)
	bookmark.DELETE("/delete",api.RepoHandler.DelBookmark)
}
