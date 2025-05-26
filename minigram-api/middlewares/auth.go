package middlewares

import (
	"fmt"
	"minigram-api/helpers"
	"minigram-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	var responseInfo models.ReponseInfo

	responseInfo.Status = 0

	hFunc := func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			responseInfo.Message = fmt.Sprintf("Unauthorized. %s", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseInfo)
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}

	return hFunc
}
