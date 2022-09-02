package r

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AR(a *gin.Engine) {
	ar := a.Group("/auth")

	// TODO: 회원가입
	ar.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "성공",
		})
	})

	// TODO: 로그인
}
