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
	Id        primitive.ObjectID `json:"_id"`
	StoreId   primitive.ObjectID `json:"storeId"`
	Name      string             `json:"name"`
	Price     string             `json:"price"`
	GoodsType GSE                `json:"goodsType"`
	Category  GSC                `json:"category"`
	Quantity  int                `json:"quantity"`
	Photos    []Photos           `json:"photos"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	IsDeleted bool               `json:"isDeleted"`
}
