package routers

import (
	"example/web-service-gin/pkg"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	versionRouter := app.Group(pkg.Env("APP_VERSION"))

	GlobalRouter(versionRouter)
}

func GlobalRouter(versionRouter *gin.RouterGroup) {
	UserRouters(versionRouter)
	AuthRouters(versionRouter)
}
