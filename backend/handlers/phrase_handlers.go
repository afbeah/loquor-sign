package handlers

import(
	"context"
	"net/http"
	"time"

	"loquor-sign/database"
	"loquor-sign/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func GetPhrase (c echo.Context) error {
	collection := database.DB.Collection("phrases")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao buscar frases",
		})
	}
	defer cursor.Close(ctx)

	var phrases []models.Phrase

	for cursor.Next(ctx) {
		var phrase models.Phrase

		if err := cursor.Decode(&phrase); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "erro ao decodificar frases",
			})
		}

		phrases = append(phrases, phrase)
	}

	if err := cursor.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao iterar frase",
		})
	}

	return c.JSON(http.StatusOK, phrases)
}

func CreatePhrase(c echo.Context) error {
	var req models.PhraseRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	userIDFromToken := c.Get("user_id").(string)

	userID, err := primitive.ObjectIDFromHex(userIDFromToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "user_id inválido",
		})
	}

	var symbols []primitive.ObjectID
	for _, s := range req.Symbols{
		objID, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "symbol_id inválido",
			})
		}

		symbols = append(symbols, objID)
	}

	phrase := models.Phrase{
		UserID: userID,
		Symbols: symbols,
		CreatedAt: time.Now(),
	}

	collection := database.DB.Collection("phrases")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, phrase)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao salvar frases",
		})
	}

	phrase.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, phrase)

}

func UpdatePhrase (c echo.Context) error{
	id := c.Param("id")

	var updatePhrase models.Phrase

	if err := c.Bind(&updatePhrase); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	collection := database.DB.Collection("phrases")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"user_id": updatePhrase.UserID,
			"symbols": updatePhrase.Symbols,
			"created_at": updatePhrase.CreatedAt,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao atualizar frases",
		})
	}

	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "frase não encontrada",
		})
	}

	updatePhrase.ID = objectID
	return c.JSON(http.StatusOK, updatePhrase)
}

func DeletePhrase (c echo.Context) error{
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	collection := database.DB.Collection("phrases")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao deletar frase",
		})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "frase não encontrada",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "frase deletada com sucesso",
	})
}

func ResetPhrases(c echo.Context) error {
	collection := database.DB.Collection("phrases")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao limpar frases",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message":"frases resetadas com sucesso",
	})

}