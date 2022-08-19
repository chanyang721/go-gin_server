package routers

import "github.com/gin-gonic/gin"

func AuthRouters(versionRouter *gin.RouterGroup) {
	authRouter := versionRouter.Group("/auth")
	{
		authRouter.GET("/:id")
	}

}
