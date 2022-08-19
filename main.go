package main

import (
	"context"
	"example/web-service-gin/pkg"
	"time"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := gin.Default()

	SetupMode(pkg.Env("GO_ENV"))

	SetupDatabase(pkg.Env("MONGO_DB_URL"))

	SetupMiddlewares(app)

	SetupRouters(app)

	logrus.Fatal(app.Run(":" + pkg.Env("GO_PORT")))
}

func SetupMode(env string) {
	if env != "production" && env != "test" {
		gin.SetMode(gin.DebugMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func SetupDatabase(url string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		logrus.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		logrus.Fatal(err)
	}

	println("Successfully Database Connected")

	defer client.Disconnect(ctx)

	return client
}

func SetupMiddlewares(app *gin.Engine) *gin.Engine {

	// CORS 옵션 //
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Helmet 옵션 //
	app.Use(helmet.Default())

	// file compresstion //

	return app
}

func SetupRouters(app *gin.Engine) *gin.Engine {

	return app
}
