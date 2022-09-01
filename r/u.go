package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UR(a *gin.Engine) {
	ur := a.Group("/users")
	ur.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "성공",
		})
	})
}
