package md

import "github.com/gin-gonic/gin"

func StpMd(env string) {
	if env == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if env == "t" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
