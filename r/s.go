package r

import (
	"github.com/gin-gonic/gin"
)

func SR(a *gin.Engine) {
	sr := a.Group("/store")

	sr.GET("/")

}
