package controllers

import (
	"fmt"
	"minigram-api/helpers"
	"minigram-api/models"
	"minigram-api/repo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User
	var exist bool
	var message string
	db := repo.GetDb()

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})

		return
	}

	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&exist)

	if err != nil {
		panic(err)
	}

	if exist {
		message := fmt.Sprintf("Username %s has been registered", user.Username)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	sqlStatement := `INSERT INTO users (username, full_name, password, token, created_date) VALUES (?, ?, ?, ?, ?)`

	currentTime := time.Now()
	pwd := helpers.HashPassword(user.Password)
	token, err := helpers.GenerateJWT(user.Username, pwd)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(sqlStatement, user.Username, user.FullName, pwd, token, currentTime)

	if err != nil {
		message = fmt.Sprintf("Error %s", err)
		panic(message)
	}

	message = fmt.Sprintf("%s Success registered", user.Username)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": message})
}
