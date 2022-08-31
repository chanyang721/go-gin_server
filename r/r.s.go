package r

import (
	"ts-s/m"

	"github.com/gin-gonic/gin"
)

func SR(a *gin.Engine) {
	sr := a.Group("/store").Use(m.A())
	{
		sr.GET("/")
		sr.POST("/")
		sr.PUT("/")
		sr.DELETE("/")
	}
}
