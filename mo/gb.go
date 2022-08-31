package mo

import (
	"time"

	"go.mongodb.org/mongo-driver/internal/uuid"
)

type GB struct {
	id        uuid.UUID
	giftId    uuid.UUID
	userId    uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
