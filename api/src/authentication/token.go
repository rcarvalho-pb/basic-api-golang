package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId uint32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
		"userId":     userId,
	})

	return token.SignedString([]byte(config.Secret))
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return errors.New("invalid token")
	}

	return nil
}

func getValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(config.Secret), nil
}

func ExtractUserId(r *http.Request) (uint32, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, getValidationKey)
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userId"]), 10, 32)
		if err != nil {
			return 0, nil
		}

		return uint32(userId), nil
	}

	return 0, errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}
