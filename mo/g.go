package mo

import (
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/internal/uuid"
)

type GE string
type GC string

type G struct {
	id              uuid.UUID
	giftType        GE
	category        GC
	quantity        int
	latitude        decimal.Decimal
	longitude       decimal.Decimal
	QRcode          string
	reservationTime time.Time
	sender          uuid.UUID
	recipient       uuid.UUID
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       time.Time
}
