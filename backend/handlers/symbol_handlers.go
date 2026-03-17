package handlers

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
  "loquor-sign/models"
)

var symbols = []models.Symbol{
		{ID:"1", Name:"Água", Image:"agua.png", Category_ID:"1"},
		{ID:"2", Name:"Comer", Image:"comer.png", Category_ID:"2"},
		{ID:"3", Name:"Dormir", Image:"dormir.png", Category_ID:"3"},
	}

func GetSymbols (c echo.Context) error {
	return c.JSON(http.StatusOK, symbols)
}

func CreateSymbol (c echo.Context) error{
	var symbol models.Symbol

	if err := c.Bind(&symbol); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	symbols = append(symbols, symbol)

	return c.JSON(http.StatusCreated, symbol)
}

func UpdateSymbol (c echo.Context) error{
	id := c.Param("id")

	var updateSymbol models.Symbol

	if err := c.Bind(&updateSymbol); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	for i, symbol := range symbols{
		if symbol.ID == id {
			symbols[i] = updateSymbol
			return c.JSON(http.StatusOK, updateSymbol)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Símbolo não encontrado",
	})
}

func DeleteSymbol (c echo.Context) error{
	id := c.Param("id")

	for i, symbol := range symbols{
		if symbol.ID == id {
			symbols = append(symbols[:i], symbols[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{
				"message": "símbolo deletado com sucesso",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "símbolo não encontrado",
	})
}