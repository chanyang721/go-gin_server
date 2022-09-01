package mo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UE string // 유저 유형

const (
	ADMIN   UE = "ADMIN"
	PARTNER UE = "PARTNER"
	NORMAL  UE = "NORMAL"
)

type U struct {
	Id        primitive.ObjectID
	PassId    string
	Email     string
	UserType  UE
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
