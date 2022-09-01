package mo

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GE string
type GC string

type G struct {
	Id              primitive.ObjectID
	Sender          primitive.ObjectID
	Recipient       primitive.ObjectID
	GiftType        GE
	Category        GC
	QRcode          string
	Use             bool
	latitude        decimal.Decimal
	longitude       decimal.Decimal
	ReservationTime time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
