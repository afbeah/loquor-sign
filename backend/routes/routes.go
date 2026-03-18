package routes

import (
	"github.com/labstack/echo/v4"
  "loquor-sign/handlers"
)

func InitRoutes(e *echo.Echo){
	e.GET("/symbols", handlers.GetSymbols)
	e.GET("/categories", handlers.GetCategory)
	e.POST("/symbols", handlers.CreateSymbol)
	e.POST("/categories", handlers.CreateCategory)
	e.PUT("/symbols/:id", handlers.UpdateSymbol)
	e.PUT("/categories/:id", handlers.UpdateCategory)
	e.DELETE("/symbols/:id", handlers.DeleteSymbol)
	e.DELETE("/categories/:id", handlers.DeleteCategory)
}
