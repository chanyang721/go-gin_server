package main

import (
	"example/web-service-gin/config/database"
	"example/web-service-gin/middlewares"
	"example/web-service-gin/pkg"
	"example/web-service-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	SetupMode(pkg.Env("GO_ENV"))

	database.SetupDatabase(pkg.Env("MONGO_DB_URL"))

	middlewares.SetupDefaultMiddleware(app)

	routers.SetupRouters(app)

	app.Run(":" + pkg.Env("GO_PORT"))
}

func SetupMode(env string) {
	if env == "development" {
		gin.SetMode(gin.DebugMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
