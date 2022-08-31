package d

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func StpMgD(mgDUl string) *mongo.Database {
	c := MgDCn(mgDUl)
	d := CMgD(c)

	return d
}

// func SetupPostgresDatabase(postgresUrl string) *sql.DB {
// 	pgDb := CreatePostgresDatabase(postgresUrl)

// 	return pgDb
// }
