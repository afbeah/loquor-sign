package main

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"loquor-sign/routes"
	"loquor-sign/database"
)

func main() {

	database.Connect()
	e := echo.New()

	routes.InitRoutes(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "api funcionando",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}