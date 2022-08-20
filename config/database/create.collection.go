package database

import (
	"example/web-service-gin/pkg"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDatabase(client *mongo.Client) *mongo.Database {
	db := client.Database(pkg.Env("Mongo_DB_Name"))

	return db
}
