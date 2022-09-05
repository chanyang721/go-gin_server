package mo

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GE string
type GC string

type G struct {
	Id              primitive.ObjectID `json:"_id"`
	Sender          U                  `json:"sender"`
	Receiver        U                  `json:"receiver,omitempty"`
	EventId         primitive.ObjectID `json:"eventId"`
	Goods           Gs                 `json:"goods"`
	GiftType        GE                 `json:"giftType"`
	Category        GC                 `json:"category"`
	QrCode          string             `json:"qrCode,omitempty"`
	Use             bool               `json:"use"`
	Confirm         bool               `json:"confirm"`
	Latitude        decimal.Decimal    `json:"latitude,omitempty"`
	Longitude       decimal.Decimal    `json:"longitude,omitempty"`
	ReservationTime time.Time          `json:"reservationTime,omitempty"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	IsDeleted       bool               `json:"isDeleted"`
}
