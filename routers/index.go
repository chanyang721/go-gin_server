package routers

import "github.com/gin-gonic/gin"

func GlobalRouter(versionRouter *gin.RouterGroup) {
	UserRouters(versionRouter)
	AuthRouter(versionRouter)
}
