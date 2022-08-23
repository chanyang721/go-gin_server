package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoDatabaseConnection(mongoDbUrl string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	println("Successfully Database Connected")

	defer client.Disconnect(ctx)

	return client
}

func PostgresDatabaseConnection(pgUrl string) *sql.DB {
	pgBb, err := sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatal(err)
	}

	return pgBb
}
