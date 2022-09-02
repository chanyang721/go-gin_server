package r

import (
	"github.com/gin-gonic/gin"
)

func SR(a *gin.Engine) {
	sr := a.Group("/store")

	// TODO: 상점 정보 조회
	sr.GET("/", func(ctx *gin.Context) {})

	// TODO: 플랫폼의 상점 리스트 조회
	sr.GET("/list", func(ctx *gin.Context) {})

}
