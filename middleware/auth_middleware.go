package middleware

import (
	"api-alemao/dto/responses"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secretJWT = []byte("chave-secreta")

func ValidationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token vazio (middleware)"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ID inválido (middleware)"})
		}

		tokenSTR := parts[1]

		token, err := jwt.Parse(tokenSTR, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("método de assinatura inesperado (middleware)")
			}
			return secretJWT, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Token no formato inválido (middleware)"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Falha ao ler o json (middleware)"})
		}

		userIDSTR, ok := claims["sub"].(string)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "ID não está presente no token"})
		}

		userIDUint, err := strconv.ParseUint(userIDSTR, 10, 64)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "ID malformado no token (middleware)"})
		}

		userID := uint(userIDUint)

		c.Set("user", userID)

		return next(c)
	}
}
