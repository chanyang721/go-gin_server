package r

import (
	"github.com/gin-gonic/gin"
)

func GsR(a *gin.Engine) {
	gsr := a.Group("/goods")

	// TODO: 상품 정보 조회
	gsr.GET("/", func(ctx *gin.Context) {})

	//TODO: 플랫폼 상점의 상품 리스트 조회
	gsr.GET("/list", func(ctx *gin.Context) {})

}
