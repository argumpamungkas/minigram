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

	_ = ctx.ShouldBind(&posting)

	// get isi bearer token
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	username := userData["username"].(string)

	photo, err := ctx.FormFile("photo")
	if err != nil {
		log.Println("ERROR PHOTO", err)
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	stringRandom := make([]byte, rand.Intn(91)+10)
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
				respInfo.Message = "failed create new folder"
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, respInfo)
				return
			}
		}

		fileName := string(stringRandom) + "." + ext
		filePath := folder + "/" + fileName
		ctx.SaveUploadedFile(photo, filePath)

		// urlPhoto := "http://" + ctx.Request.Host + "/img/" + username + "/" + fileName
		// photoName := username + "/" + fileName
		posting.Photo = fileName
	} else {
		respInfo.Message = "file is not Image"
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
	respInfo.Message = "photo success upload"

	ctx.JSON(http.StatusCreated, respInfo)

}

func DeletePosting(ctx *gin.Context) {
	var respInfo models.ReponseInfo
	var posting models.Posting

	db := repo.GetDb()

	respInfo.Status = 0

	postingId := ctx.Param("postingId")

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	username := userData["username"].(string)
	userId := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ? AND user_id = ?", postingId, userId).First(&posting).Error
	if err != nil {
		log.Println("SELECT POSTING")
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	folder := "./assets/" + username + "/" + posting.Photo

	log.Println("FOLDER => ", folder)
	err = os.Remove(folder)
	if err != nil {
		log.Println("DELETE FILE")
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	err = db.Debug().Where("id = ?", postingId).Delete(&posting).Error
	if err != nil {
		log.Println("SELECT POSTING")
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	respInfo.Status = 1
	respInfo.Message = "Photo success Deleted"

	ctx.JSON(http.StatusOK, respInfo)
}

func GetPostingById(ctx *gin.Context) {
	var respInfo models.ReponseInfo
	var responsePosting models.ReponsePostingById
	var posting models.Posting

	db := repo.GetDb()

	respInfo.Status = 0

	postingId := ctx.Param("postingId")

	err := db.Debug().Preload("User").Where("id = ?", postingId).First(&posting).Error
	if err != nil {
		log.Println("SELECT POSTING")
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	responsePosting.Status = 1
	responsePosting.Message = "Success"
	responsePosting.Data.UserId = posting.UserId
	responsePosting.Data.Username = posting.User.Username
	responsePosting.Data.Avatar = posting.User.Avatar
	responsePosting.Data.Photo = posting.Photo
	responsePosting.Data.Caption = posting.Caption
	responsePosting.Data.Likes = 0
	responsePosting.Data.Comments = 0

	ctx.JSON(http.StatusOK, responsePosting)
}

func GetPostingAll(ctx *gin.Context) {
	var respInfo models.ReponseInfo
	var responsePosting models.ReponsePostingAll
	var postings []models.Posting

	db := repo.GetDb()

	respInfo.Status = 0

	err := db.Debug().Preload("User").Order("RANDOM()").Find(&postings).Error
	if err != nil {
		log.Println("SELECT POSTING")
		respInfo.Message = err.Error()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respInfo)
		return
	}

	responsePosting.Status = 1
	responsePosting.Message = "success"
	responsePosting.Data = []models.ResponsePosting{} // untuk mengindari null jika tidak ada datanya -> []

	for _, value := range postings {
		rp := models.ResponsePosting{}
		rp.UserId = value.UserId
		rp.Username = value.User.Username
		rp.Avatar = value.User.Avatar
		rp.Photo = value.Photo
		rp.Caption = value.Caption
		rp.Likes = 0
		rp.Comments = 0

		responsePosting.Data = append(responsePosting.Data, rp)
	}

	ctx.JSON(http.StatusOK, responsePosting)
}
