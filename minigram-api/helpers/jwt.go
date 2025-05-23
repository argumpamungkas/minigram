package helpers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = "key123Secret"

func GenerateJWT(username string, email string) (res string, err error) {

	currentTime := time.Now().Unix()

	claims := jwt.MapClaims{
		"email":    email,
		"username": username,
		"time":     currentTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("JWT signing error:", err)
		return
	}
	res = signedToken

	return
}
