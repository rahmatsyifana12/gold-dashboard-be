package main

import (
	"gold-dashboard-be/internal/apps"

	"github.com/labstack/echo/v4"
)

type Module struct{}

func (m *Module) New(app *echo.Echo) {
	ioc := apps.NewIOC()

	r := NewRoute(app, ioc)
	r.Init()
}