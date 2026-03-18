package handlers

import (
	"context"
	"net/http"
	"time"

	"loquor-sign/database"
	"loquor-sign/models"
	
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var symbols = []models.Symbol{
		{ID:"1", Name:"Água", Image:"agua.png", Category_ID:"1"},
		{ID:"2", Name:"Comer", Image:"comer.png", Category_ID:"2"},
		{ID:"3", Name:"Dormir", Image:"dormir.png", Category_ID:"3"},
	}

func GetSymbols(c echo.Context) error {
	collection := database.DB.Collection("symbols")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao buscar símbolo",
		})
	}
	defer cursor.Close(ctx)

	var symbols []models.Symbol

	for cursor.Next(ctx) {
		var symbol models.Symbol

		if err := cursor.Decode(&symbol); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "erro ao decodificar símbolo",
			})
		}

		symbols = append(symbols, symbol)
	}

	return c.JSON(http.StatusOK, symbols)
}

func CreateSymbol (c echo.Context) error{
	var symbol models.Symbol

	if err := c.Bind(&symbol); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	collection := database.DB.Collection("symbols")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, symbol)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao salvar no banco",
		})
	}

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

	collection := database.DB.Collection("symbols")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name": updateSymbol.Name,
			"image": updateSymbol.Image,
			"category_id": updateSymbol.Category_ID,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao atualizar símbolo",
		})
	}

	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "símbolo não encontrado",
		})
	}

	return c.JSON(http.StatusOK, updateSymbol)
}

func DeleteSymbol (c echo.Context) error{
	id := c.Param("id")

	collection := database.DB.Collection("symbols")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao deletar símbolo",
		})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "símbolo não encontrado",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "símbolo deletado com sucesso",
	})
}