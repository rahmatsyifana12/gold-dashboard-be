package main

import (
	"gold-dashboard-be/internal/apps/rest/handlers"
	"gold-dashboard-be/internal/apps/rest/middlewares"
	"gold-dashboard-be/internal/constants"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type Route struct {
	app        *echo.Echo
	router     *echo.Group
	controller *handlers.Controller
}

func NewRoute(app *echo.Echo, ioc di.Container) *Route {
	return &Route{
		app:        app,
		router:     app.Group(""),
		controller: ioc.Get(constants.Controller).(*handlers.Controller),
	}
}

func (r *Route) Init() {
	r.Ping()
	r.Auth()
	r.User()
}

func (r *Route) Ping() {
	r.app.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
}

func (r *Route) User() {
	user := r.router.Group("users")
	user.POST("", r.controller.User.CreateUser)
	user.GET("/:id", r.controller.User.GetUserByID, middlewares.AuthMiddleware)
	user.PATCH("/:id", r.controller.User.UpdateUser, middlewares.AuthMiddleware)
	user.DELETE("/:id", r.controller.User.DeleteUser, middlewares.AuthMiddleware)
}

func (r *Route) Auth() {
	auth := r.router.Group("auth")
	auth.POST("/login", r.controller.Auth.Login)
	// auth.POST("/logout", r.controller.Auth.Logout, middlewares.AuthMiddleware)
}
