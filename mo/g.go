package mo

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GE string
type GC string

type G struct {
	Id              primitive.ObjectID `json:"id"`
	Sender          primitive.ObjectID `json:"sender"`
	Receiver        primitive.ObjectID `json:"receiver"`
	GiftType        GE                 `json:"giftType"`
	Category        GC                 `json:"category"`
	QrCode          string             `json:"qrCode"`
	Use             bool               `json:"use"`
	Confirm         bool               `json:"confirm"`
	latitude        decimal.Decimal    `json:"latitude"`
	longitude       decimal.Decimal    `json:"longitude"`
	ReservationTime time.Time          `json:"reservationTime"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"UpdatedAt"`
	DeletedAt       time.Time          `json:"deletedAt"`
}
