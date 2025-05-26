package helpers

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = "key123Secret"

func GenerateJWT(id uint, username string, email string) (res string, err error) {

	currentTime := time.Now().Unix()

	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
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

func VerifyToken(ctx *gin.Context) (res interface{}, err error) {
	header := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(header, "Bearer")

	if !bearer {
		err = errors.New("invalid header")
		return
	}

	keyFunc := func(j *jwt.Token) (res interface{}, err error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.New("invalid method header")
			return
		}

		res, err = []byte(key), nil
		return
	}

	log.Println("str token ", header)

	strToken := strings.Split(header, " ")[1]
	token, err := jwt.Parse(strToken, keyFunc)

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		err = errors.New("invalid token")
		return
	}

	res = token.Claims.(jwt.MapClaims)
	return
}
