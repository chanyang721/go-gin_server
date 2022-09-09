package main

import (
	"ts-s/cfg/md"
	"ts-s/d"
	"ts-s/m"
	"ts-s/p/e"
	"ts-s/r"

	"github.com/gin-gonic/gin"
)

func main() {
	a := gin.Default()

	md.StpMd(e.G("GO_ENV"))

	d.StpMgD(e.G("MONGO_DB_URL"))

	m.StpM(a)

	r.StpR(a)

	a.Run(":" + e.G("GO_PORT"))
}
