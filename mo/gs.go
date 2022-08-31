package mo

import (
	"time"

	"go.mongodb.org/mongo-driver/internal/uuid"
)

type GSE string // 상품 유형
type GSC string // 상품 카테고리

const (
	TICKET GE = "TICKET"
	REAL   GE = "REAL"
	EVENT  GE = "EVENT"
)

const ()

type GS struct {
	id        uuid.UUID
	storeId   uuid.UUID
	name      string
	price     string
	goodsType GSE // 상품 유형
	category  GSC // 상품 카테고리
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
