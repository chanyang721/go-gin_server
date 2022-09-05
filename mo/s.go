package mo

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SE string // 상점 유형
type SC string // 상점 카테고리
type PE string // 플랫폼 유형

const (
	PLATFORM SE = "PLATFORM"
	TOSPACE  SE = "TOSPACE"
)

const ()

type S struct {
	Id            primitive.ObjectID `json:"_id"`
	Name          string             `json:"name" binding:"required"`
	Phone         string             `json:"phone" binding:"required"`
	StoreType     SE                 `json:"storeType,omitempty" binding:"required"`
	Category      SC                 `json:"category,omitempty" binding:"required"`
	PlatformType  PE                 `json:"platformType,omitempty"`
	ShopName      string             `json:"shopName" binding:"required"`
	ShopAddress   string             `json:"shopAddress" binding:"required"`
	CompanyNumber string             `json:"companyNumber" binding:"required"`
	Latitude      decimal.Decimal    `json:"latitude" binding:"required"`
	Longitude     decimal.Decimal    `json:"longitude" binding:"required"`
	Radius        int                `json:"radius,omitempty"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	IsDeleted     bool               `json:"isDeleted"`
}
