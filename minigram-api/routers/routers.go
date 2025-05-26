package routers

import (
	"minigram-api/controllers"
	"minigram-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.Login)
	}

	router.Static("/img", "./assets")
	posting := router.Group("/posting")
	{
		posting.Use(middlewares.Auth())

		posting.POST("/post", controllers.CreatePosting)
	}

	return router
}
