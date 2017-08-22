package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/beewit/beekit/utils"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middleware.Gzip())

	e.Static("/page", "page")
	e.Static("/static", "static")

	e.File("/","page/index.html")

	utils.Open("http://127.0.0.1:8080")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
