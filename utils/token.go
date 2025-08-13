package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretJWT = []byte("chave-secreta")

func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.FormatUint(uint64(id), 10),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenSTR, err := token.SignedString(secretJWT)
	if err != nil {
		return "", fmt.Errorf("falha ao gerar token: %v", err)
	}
	return tokenSTR, nil
}
