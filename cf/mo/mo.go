package mo

import "github.com/gin-gonic/gin"

func StpM(env string) {
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if env == "t" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
