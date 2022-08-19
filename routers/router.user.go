package routers

import (
	"example/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouters(versionRouter *gin.RouterGroup) {
	userRouter := versionRouter.Group("/users").Use(middlewares.Auth())
	{
		userRouter.GET("/:id")
		userRouter.POST("/")
		userRouter.PUT("/:id")
		userRouter.DELETE("/:id")
	}
}
