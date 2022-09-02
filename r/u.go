package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UR(a *gin.Engine) {
	ur := a.Group("/users")

	// TODO: 사용자 정보 조회
	ur.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "성공",
		})
	})

	// TODO: 사용자 정보 수정

}
