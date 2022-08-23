package main

import (
	"example/web-service-gin/config/database"
	"example/web-service-gin/config/mode"
	"example/web-service-gin/middlewares"
	"example/web-service-gin/pkg"
	"example/web-service-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	mode.SetupMode(pkg.Env("GO_ENV"))

	database.SetupMongoDatabase(pkg.Env("MONGO_DB_URL"))

	database.SetupPostgresDatabase(pkg.Env("POSTGRESQL_DB_URL"))

	middlewares.SetupDefaultMiddleware(app)

	routers.SetupRouters(app)

	app.Run(":" + pkg.Env("GO_PORT"))
}
