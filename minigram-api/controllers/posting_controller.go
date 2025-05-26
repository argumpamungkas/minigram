package controllers

import (
	"log"
	"minigram-api/models"
	"minigram-api/repo"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/rand"
)

func CreatePosting(ctx *gin.Context) {
	var respInfo models.ReponseInfo
	var posting models.Posting

	db := repo.GetDb()

	// init status 0
	respInfo.Status = 0

	// get berdasarkan bearer token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	username := userData["username"].(string)

	// caption, err := ctx.form("caption")
	// if err != nil {
	// 	log.Println("ERROR Caption", err)
	// 	respInfo.Message = err.Error()
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
	// 	return
	// }

	photo, err := ctx.FormFile("photo")
	if err != nil {
		log.Println("ERROR PHOTO", err)
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	stringRandom := make([]byte, rand.Intn(100))
	for i := range stringRandom {
		stringRandom[i] = charset[rand.Intn(len(charset))]
	}

	ext := strings.Split(photo.Filename, ".")[1]

	log.Println("file ext ->", ext)

	if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "webp" {
		// path folder berdasarkan username
		folder := "./assets/" + username

		// Cek folder apakah sudah ada atau belum?
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			// jika belum ada foldernya maka dibuat
			err := os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				respInfo.Message = "Failed create new folder"
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, respInfo)
				return
			}
		}

		fileName := string(stringRandom) + "." + ext
		filePath := folder + "/" + fileName
		ctx.SaveUploadedFile(photo, filePath)

		// urlPhoto := "http://" + ctx.Request.Host + "/img/" + username + "/" + fileName
		photoName := username + "/" + fileName
		posting.Photo = photoName
	} else {
		respInfo.Message = "File is not Image"
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respInfo)
		return
	}

	// "http://" + ctx.Request.Host + "/img/" + username + "-" + string(stringRandom) + "." + ext
	// posting.Caption = caption.Filename
	posting.UserId = userID

	log.Println("POSTING => ", posting)

	_, err = posting.BeforeCreate()
	if err != nil {
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	err = db.Debug().Create(&posting).Error
	if err != nil {
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, respInfo)
		return
	}

	respInfo.Status = 1
	respInfo.Message = "Photo success upload"

	ctx.JSON(http.StatusCreated, respInfo)

}
