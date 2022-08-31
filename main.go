package main

import (
	"ts-s/cf/mo"
	"ts-s/d"
	"ts-s/m"
	"ts-s/p/e"
	"ts-s/r"

	"github.com/gin-gonic/gin"
)

func main() {
	a := gin.Default()

	mo.StpM(e.G("GO_ENV"))

	d.StpMgD(e.G("MONGO_DB_URL"))

	// d.SetupPostgresDatabase(pkg.Env("POSTGRESQL_DB_URL"))

	m.StpM(a)

	r.SR(a)

	a.Run(":" + e.G("GO_PORT"))
}
