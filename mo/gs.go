package mo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GSE string // 상품 유형
type GSC string // 상품 카테고리

const (
	TICKET GSE = "TICKET"
	REAL   GSE = "REAL"
	EVENT  GSE = "EVENT"
)

const ()

type Gs struct {
	Id        primitive.ObjectID
	StoreId   primitive.ObjectID
	Name      string
	Price     string
	GoodsType GSE // 상품 유형
	Category  GSC // 상품 카테고리
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
