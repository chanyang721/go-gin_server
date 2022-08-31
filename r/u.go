package r

import (
	"github.com/gin-gonic/gin"
)

func UR(a *gin.Engine) {
	ur := a.Group("/users")
	{
		ur.GET("/")
		ur.POST("/")
		ur.PUT("/")
		ur.DELETE("/")
	}
}
