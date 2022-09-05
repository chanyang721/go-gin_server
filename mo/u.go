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
	Id        primitive.ObjectID `json:"_id"`
	PassId    string             `json:"-"`
	Name      string             `json:"name" binding:"required"`
	Email     string             `json:"email,omitempty"`
	EmailAuth bool               `json:"emailAuth"`
	UserType  UE                 `json:"userType"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	IsDeleted bool               `json:"isDeleted"`
}
