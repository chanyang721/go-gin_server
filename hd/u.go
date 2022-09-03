package hd

import (
	"ts-s/mo"
	"ts-s/s"

	"github.com/gin-gonic/gin"
)

type Uhd struct {
	s s.IUS
}

func UHd(c *gin.Context) {
	i := mo.U{}
	c.ShouldBind(&i)

	// 서비스 호출

	//

	return
}
