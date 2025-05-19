package helpers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) (res string) {
	salt := 8
	password := []byte(s)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		msg := fmt.Sprintf("Error generate pwd %s", err)
		log.Println(msg)
		return
	}

	res = string(hash)
	return
}
