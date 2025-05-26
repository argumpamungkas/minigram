package controllers

import (
	"fmt"
	"log"
	"minigram-api/helpers"
	"minigram-api/models"
	"minigram-api/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User
	var responseLogin models.ReponseLogin
	var responseInfo models.ReponseInfo
	// var exist bool
	var count int64
	db := repo.GetDb()

	responseInfo.Status = 0

	if err := ctx.ShouldBindJSON(&user); err != nil {
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// CEK USERNAME
	err := db.Debug().Model(&user).Where("username = ?", user.Username).Count(&count).Error
	// err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&exist)
	if err != nil {
		responseInfo.Message = fmt.Sprintf("Error Username %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	if count > 0 {
		responseInfo.Message = fmt.Sprintf("Username %s has been registered", user.Username)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// CEK EMAIL
	err = db.Debug().Model(&user).Where("email = ?", user.Email).Count(&count).Error
	// err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", user.Email).Scan(&exist)
	if err != nil {
		responseInfo.Message = fmt.Sprintf("Error Email %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	if count > 0 {
		responseInfo.Message = fmt.Sprintf("Email %s has been registered", user.Email)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// Insert USER
	// sqlStatement := `INSERT INTO users (username, full_name, email, password, created_date) VALUES ( ?, ?, ?, ?, ?)`

	_, err = user.BeforeCreate()
	if err != nil {
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// Perintah SQL
	err = db.Debug().Create(&user).Error
	// _, err = db.Exec(sqlStatement, user.Username, user.FullName, user.Email, user.Password, currentTime)

	if err != nil {
		log.Println("EXEC", err)
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// If success
	token, err := helpers.GenerateJWT(user.Id, user.Username, user.Email)
	if err != nil {
		log.Println("Error generate token", err)
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	responseLogin.Status = 1
	responseLogin.Message = fmt.Sprintf("%s Success registered", user.Username)
	responseLogin.User.FullName = user.FullName
	responseLogin.User.Email = user.Email
	responseLogin.User.Username = user.Username
	responseLogin.User.Token = token

	ctx.JSON(http.StatusCreated, responseLogin)
}

func Login(ctx *gin.Context) {
	var user models.User
	var requsetLogin models.RequestLogin
	var responseLogin models.ReponseLogin
	var responseInfo models.ReponseInfo
	// var exist bool
	var count int64
	db := repo.GetDb()

	responseInfo.Status = 0

	if err := ctx.ShouldBindJSON(&requsetLogin); err != nil {
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// CEK USERNAME
	err := db.Debug().Model(&user).Where("username = ?", requsetLogin.Username).Count(&count).Error
	// err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", user.Username).Scan(&exist)
	if err != nil {
		responseInfo.Message = fmt.Sprintf("Error Username %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	if count == 0 {
		responseInfo.Message = fmt.Sprintf("Username %s Not found", requsetLogin.Username)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// CEK Password
	err = db.Debug().Model(&user).Select("password").Where("username = ?", requsetLogin.Username).Scan(&user.Password).Error

	// err = db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&pwd)
	if err != nil {
		responseInfo.Message = fmt.Sprintf("Error Password %s", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	comparePass := helpers.ComparePassword([]byte(user.Password), []byte(requsetLogin.Password))
	if !comparePass {
		responseInfo.Message = "Invalid password"
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// Insert USER
	// sqlStatement := `SELECT username, full_name, email, avatar FROM users WHERE username = ?`

	// Perintah SQL
	err = db.Debug().Where("username = ?", requsetLogin.Username).Take(&user).Error
	// err = db.QueryRow(sqlStatement, user.Username).Scan(&responseLogin.User.Username, &responseLogin.User.FullName, &responseLogin.User.Email, &responseLogin.User.Avatar)

	if err != nil {
		log.Println("GET DATA", err)
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	// If success
	token, err := helpers.GenerateJWT(user.Id, user.Username, user.Email)
	if err != nil {
		log.Println("Error generate token", err)
		responseInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseInfo)
		return
	}

	log.Println(ctx.Request.Host)

	responseLogin.Status = 1
	responseLogin.Message = "User founded"
	responseLogin.User.Username = user.Username
	responseLogin.User.Email = user.Email
	responseLogin.User.FullName = user.FullName
	responseLogin.User.Token = token

	ctx.JSON(http.StatusCreated, responseLogin)
}
