package routers

import (
	"example/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouters(versionRouter *gin.RouterGroup) {
	userRouters := versionRouter.Group("/users").Use(middlewares.Auth())
	{
		userRouters.GET("/:id")
		userRouters.POST("/")
		userRouters.PUT("/:id")
		userRouters.DELETE("/:id")
	}
}
