package cn

import (
	"context"
	"fmt"
	"time"
	"ts-s/p/e"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MgDCn(mgDU string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mgDU))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	println("Successfully connected and pinged")

	d := client.Database(e.G("Mongo_DB_Name"))
	u := d.Collection("users")
	d.Collection("stores")
	d.Collection("gifts")
	d.Collection("goods")

	result, _ := u.InsertOne(ctx, bson.D{})
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	return client
}
