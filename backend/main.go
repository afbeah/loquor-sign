package main

import (
	"log"
	"net/http"

	"loquor-sign/routes"
	"loquor-sign/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	database.Connect()

	e := echo.New()
	e.Use(middleware.CORS())

	routes.InitRoutes(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "api funcionando",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}