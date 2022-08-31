package r

import (
	"ts-s/m"

	"github.com/gin-gonic/gin"
)

func GR(a *gin.Engine) {
	gr := a.Group("/gift").Use(m.A())
	{
		gr.GET("/")
		gr.POST("/")
		gr.PUT("/")
		gr.DELETE("/")
	}
}
