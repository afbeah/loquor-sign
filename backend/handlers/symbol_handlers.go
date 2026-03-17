package handlers

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
  "loquor-sign/models"
)

func GetSymbols (c echo.Context) error {
	
	symbols := []models.Symbol{
		{ID:"1", Name:"Água", Image:"agua.png", Category_ID:"1"},
		{ID:"2", Name:"Comer", Image:"comer.png", Category_ID:"2"},
		{ID:"3", Name:"Dormir", Image:"dormir.png", Category_ID:"3"},
	}

	return c.JSON(http.StatusOK, symbols)

}