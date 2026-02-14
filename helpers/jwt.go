package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string
var jwtSecretArrOfByte []byte

func InitJWT() {
	jwtSecret = os.Getenv("JWT_SECRET_KEY")
	jwtSecretArrOfByte = []byte(jwtSecret)

	// log.Println("JWT secret:", jwtSecret)
}

type AuthPayload struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Exp   uint   `json:"exp"`
}

type AuthTokenClaims struct {
	ID    uint
	Name  string
	Email string
}

func CreateAuthToken(claims AuthTokenClaims) (string, error) {
	log.Println(jwtSecret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    claims.ID,
		"name":  claims.Name,
		"email": claims.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	// log.Println("Token claim", tokenClaims)

	authToken, err := tokenClaims.SignedString(jwtSecretArrOfByte)
	if err != nil {
		log.Println("Error creating auth token:", err)
		return "", err
	}

	return authToken, nil
}

func VerifyAuthToken(strAuthToken string) (*jwt.Token, error) {
	authToken, err := jwt.Parse(strAuthToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return jwtSecretArrOfByte, nil
	})

	// log.Println(err)

	return authToken, err
}
