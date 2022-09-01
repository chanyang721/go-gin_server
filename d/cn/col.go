package cn

import (
	"ts-s/p/e"

	"go.mongodb.org/mongo-driver/mongo"
)

func CMgD(c *mongo.Client) *mongo.Database {
	d := c.Database(e.G("Mongo_DB_Name"))
	return d
}

// func CPgD(pgU string) *sql.DB {
// 	d := PgDCn(pgU)

// 	return d
// }
