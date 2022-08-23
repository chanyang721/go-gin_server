package database

import (
	"database/sql"
	"example/web-service-gin/pkg"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMongoDatabase(client *mongo.Client) *mongo.Database {
	db := client.Database(pkg.Env("Mongo_DB_Name"))
	return db
}

func CreatePostgresDatabase(pgUrl string) *sql.DB {
	db := PostgresDatabaseConnection(pgUrl)

	return db
}
