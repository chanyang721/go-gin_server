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
	Id             primitive.ObjectID `json:"id"`
	UserId         primitive.ObjectID `json:"userId"`
	Name           string             `json:"name"`
	Phone          string             `json:"phone"`
	StoreType      SE                 `json:"storeType"`
	Category       SC                 `json:"category"`
	PlatformType   PE                 `json:"platformType"`
	ShopName       string             `json:"shopName"`
	ShopAddress    string             `json:"shopAddress"`
	BusinessNumber string             `json:"businessNumber"`
	Latitude       decimal.Decimal    `json:"latitude"`
	Longitude      decimal.Decimal    `json:"longitude"`
	Radius         int                `json:"radius"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	IsDeleted      bool               `json:"IsDeleted"`
}
