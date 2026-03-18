package routes

import (
	"github.com/labstack/echo/v4"
  "loquor-sign/handlers"
)

func InitRoutes(e *echo.Echo){
	e.GET("/symbols", handlers.GetSymbols)
	e.GET("/categories", handlers.GetCategory)
	e.GET("/users", handlers.GetUser)
	e.POST("/symbols", handlers.CreateSymbol)
	e.POST("/categories", handlers.CreateCategory)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/symbols/:id", handlers.UpdateSymbol)
	e.PUT("/categories/:id", handlers.UpdateCategory)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/symbols/:id", handlers.DeleteSymbol)
	e.DELETE("/categories/:id", handlers.DeleteCategory)
	e.DELETE("/users/:id", handlers.DeleteUser)
}
