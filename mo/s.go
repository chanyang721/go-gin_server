package mo

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/internal/uuid"
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
	id             uuid.UUID
	userId         uuid.UUID
	name           string
	phone          string
	storeType      SE
	category       SC
	platformType   PE
	shopName       string
	shopAddress    string
	businessNumber string
	latitude       decimal.Decimal
	longitude      decimal.Decimal
	radius         int
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      time.Time
}
