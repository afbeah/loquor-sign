package middleware

import (
	"net/http"
	"strings"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "token ausente",
			})
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "token inválido",
			})
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "token inválido",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])

		return next(c)
	
	}

}