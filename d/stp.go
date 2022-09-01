package d

import (
	"ts-s/d/cn"

	"go.mongodb.org/mongo-driver/mongo"
)

func StpMgD(mgDUl string) *mongo.Database {
	c := cn.MgDCn(mgDUl)
	d := cn.CMgD(c)

	return d
}

// func StpPgD(pgUl string) *sql.DB {
// 	d := cn.CPgD(pgUl)

// 	return d
// }
