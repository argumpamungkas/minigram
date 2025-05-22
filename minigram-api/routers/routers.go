package routers

import (
	"minigram-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	return router
}
