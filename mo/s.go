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
	Id             primitive.ObjectID
	UserId         primitive.ObjectID
	Name           string
	Phone          string
	StoreType      SE
	Category       SC
	PlatformType   PE
	ShopName       string
	ShopAddress    string
	BusinessNumber string
	Latitude       decimal.Decimal
	Longitude      decimal.Decimal
	Radius         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
