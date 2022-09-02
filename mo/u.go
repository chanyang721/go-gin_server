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
	Id        primitive.ObjectID `json:"id"`
	PassId    string             `json:"passId"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	UserType  UE                 `json:"userType"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	IsDeleted bool               `json:"IsDeleted"`
}
