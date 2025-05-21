package controllers

import (
	"fmt"
	"log"
	"minigram-api/models"
	"minigram-api/repo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User
	var responseInsert models.ReponseInsert
	var exist bool
	db := repo.GetDb()

	responseInsert.Status = 0

	if err := ctx.ShouldBindJSON(&user); err != nil {
		responseInsert.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	// CEK USERNAME
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&exist)
	if err != nil {
		responseInsert.Message = fmt.Sprintf("Error Username %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	if exist {
		responseInsert.Message = fmt.Sprintf("Username %s has been registered", user.Username)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	// CEK EMAIL
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.Email).Scan(&exist)
	if err != nil {
		responseInsert.Message = fmt.Sprintf("Error Email %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	if exist {
		responseInsert.Message = fmt.Sprintf("Email %s has been registered", user.Email)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	// Insert USER
	sqlStatement := `INSERT INTO users (username, full_name, email, password, token, created_date) VALUES (?, ?, ?, ?, ?, ?)`

	_, err = user.BeforeCreate()
	if err != nil {
		responseInsert.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	currentTime := time.Now()

	// Perintah SQL
	_, err = db.Exec(sqlStatement, user.Username, user.FullName, user.Email, user.Password, user.Token, currentTime)

	if err != nil {
		log.Println("EXEC", err)
		responseInsert.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInsert)
		return
	}

	// If success
	responseInsert.Status = 1
	responseInsert.Message = fmt.Sprintf("%s Success registered", user.Username)

	ctx.JSON(http.StatusCreated, responseInsert)
}
