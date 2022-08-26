package routers

import (
	"example/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func GiftRouters(versionRouter *gin.RouterGroup) {
	giftRouter := versionRouter.Group("/gift").Use(middlewares.Auth())
	{
		giftRouter.GET("/")
		giftRouter.POST("/")
		giftRouter.PUT("/")
		giftRouter.DELETE("/")
	}
}
