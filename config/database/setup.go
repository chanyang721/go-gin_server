package database

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

func SetupMongoDatabase(mongoDbUrl string) *mongo.Database {
	client := MongoDatabaseConnection(mongoDbUrl)
	db := CreateMongoDatabase(client)

	return db
}

func SetupPostgresDatabase(postgresUrl string) *sql.DB {
	pgDb := CreatePostgresDatabase(postgresUrl)

	return pgDb
}
