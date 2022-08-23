package mode

import "github.com/gin-gonic/gin"

func SetupMode(env string) {
	if env == "development" {
		gin.SetMode(gin.DebugMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
