package main

import (
	. "github.com/johnazedo/playmovie-bff/src/home"
	"github.com/labstack/echo/v4"
)

func main() {
	home := NewHomeController()

	// Server
	e := echo.New()
	e.GET("/home", home.GetHome)
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
