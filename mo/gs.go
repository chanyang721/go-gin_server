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
	Id        primitive.ObjectID `json:"id"`
	StoreId   primitive.ObjectID `json:"storeId"`
	Name      string             `json:"name"`
	Price     string             `json:"price"`
	GoodsType GSE                `json:"goodsType"` // 상품 유형
	Category  GSC                `json:"category"`  // 상품 카테고리
	Quantity  int                `json:"quantity"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"UpdatedAt"`
	DeletedAt time.Time          `json:"deletedAt"`
}
