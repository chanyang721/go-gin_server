package mo

import "go.mongodb.org/mongo-driver/bson/primitive"

type PtE string

const (
	IMAGE PtE = "IMAGE"
	AR    PtE = "AR"
)

type Photos struct {
	Id        primitive.ObjectID
	GoodsId   primitive.ObjectID
	PhotoType PE
	Url       string
}
