package routers

import (
	"example/web-service-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func StoreRouters(versionRouter *gin.RouterGroup) {
	storeRouter := versionRouter.Group("/store").Use(middlewares.Auth())
	{
		storeRouter.GET("/")
		storeRouter.POST("/")
		storeRouter.PUT("/")
		storeRouter.DELETE("/")
	}
}
