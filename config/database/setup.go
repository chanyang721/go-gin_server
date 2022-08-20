package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupDatabase(mongoDbUrl string) *mongo.Database {
	client := DatabaseConnection(mongoDbUrl)

	db := CreateDatabase(client)

	return db
}
