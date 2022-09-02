package r

import (
	"github.com/gin-gonic/gin"
)

func StpR(a *gin.Engine) {
	UR(a)
	GR(a)
	GsR(a)
	SR(a)
}

//TODO: Request 객체 구조체 생성
//TODO: Response 객체 구조체 생성
//TODO: Error 핸들링 객체 구조체 생성
//TODO: 유효성 검사 함수 생성
