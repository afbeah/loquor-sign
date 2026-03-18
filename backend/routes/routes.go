package routes

import (
	"github.com/labstack/echo/v4"
  "loquor-sign/handlers"
)

func InitRoutes(e *echo.Echo){
	e.GET("/symbols", handlers.GetSymbols)
	e.POST("/symbols", handlers.CreateSymbol)
	e.PUT("/symbols/:id", handlers.UpdateSymbol)
	e.DELETE("/symbols/:id", handlers.DeleteSymbol)
}
