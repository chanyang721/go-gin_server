package r

import "github.com/gin-gonic/gin"

func AR(a *gin.Engine) {
	ar := a.Group("/auth")
	{
		ar.GET("/:id")
	}
}
