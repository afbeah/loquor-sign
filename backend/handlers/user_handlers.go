package handlers

import(
	"os"
	"context"
	"net/http"
	"time"

	"loquor-sign/database"
	"loquor-sign/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/golang-jwt/jwt/v5"

)

func GetUser (c echo.Context) error {
	collection := database.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao buscar usuário",
		})
	}
	defer cursor.Close(ctx)

	var users []models.User

	for cursor.Next(ctx) {
		var user models.User

		if err := cursor.Decode(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "erro ao decodificar usuário",
			})
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao iterar usuário",
		})
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao criptografar senha",
		})
	}

	user.Password = string(hashedPassword)

	collection := database.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao salvar usuário",
		})
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	response := models.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusCreated, response)

}

func UpdateUser (c echo.Context) error{
	id := c.Param("id")

	var updateUser models.User

	if err := c.Bind(&updateUser); err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	collection := database.DB.Collection("users")

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
			"name": updateUser.Name,
			"email": updateUser.Email,
			"password": updateUser.Password,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao atualizar usuário",
		})
	}

	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "usuário não encontrado",
		})
	}

	updateUser.ID = objectID
	updateUser.Password = ""
	return c.JSON(http.StatusOK, updateUser)
}

func Login(c echo.Context) error {
	var loginData struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	collection := database.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	err := collection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil{
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "usuário ou senha inválida",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "usuário ou senha inválida",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao gerar token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": models.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
		},
		"token": tokenString,
	})

}

func DeleteUser (c echo.Context) error{
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	collection := database.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "erro ao deletar usuário",
		})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "usuário não encontrado",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "usuário deletado com sucesso",
	})
}