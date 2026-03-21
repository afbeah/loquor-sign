package routes

import (
	"github.com/labstack/echo/v4"
  "loquor-sign/handlers"
)

func InitRoutes(e *echo.Echo){
	e.GET("/symbols", handlers.GetSymbols)
	e.GET("/categories", handlers.GetCategory)
	e.GET("/users", handlers.GetUser)
	e.GET("/phrases", handlers.GetPhrase)

	e.POST("/symbols", handlers.CreateSymbol)
	e.POST("/categories", handlers.CreateCategory)
	e.POST("/users", handlers.CreateUser)
	e.POST("/login", handlers.Login)
	e.POST("/phrases", handlers.CreatePhrase)

	e.PUT("/symbols/:id", handlers.UpdateSymbol)
	e.PUT("/categories/:id", handlers.UpdateCategory)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.PUT("/phrases/:id", handlers.UpdatePhrase)

	e.DELETE("/symbols/:id", handlers.DeleteSymbol)
	e.DELETE("/categories/:id", handlers.DeleteCategory)
	e.DELETE("/users/:id", handlers.DeleteUser)
	e.DELETE("/phrases/:id", handlers.DeletePhrase)
}
