package d

import (
	"ts-s/d/cn"

	"go.mongodb.org/mongo-driver/mongo"
)

func StpMgD(mgDUl string) *mongo.Client {
	c := cn.MgDCn(mgDUl)
	// d := cn.CMgD(c, ctx)

	return c
}

// func StpPgD(pgUl string) *sql.DB {
// 	d := cn.CPgD(pgUl)

// 	return d
// }
