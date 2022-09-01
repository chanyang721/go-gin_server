package r

import (
	"github.com/gin-gonic/gin"
)

func GsR(a *gin.Engine) {
	gsr := a.Group("/goods")

	gsr.GET("/")

}
