package helpers

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var key = "key123Secret"

func GenerateJWT(username string, pwd string) (res string, err error) {

	claims := jwt.MapClaims{
		"pwd":      pwd,
		"username": username,
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
