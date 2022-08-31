package mo

import (
	"time"

	"go.mongodb.org/mongo-driver/internal/uuid"
)

type UE string // 유저 유형

const (
	ADMIN   UE = "ADMIN"
	PARTNER UE = "PARTNER"
	NORMAL  UE = "NORMAL"
)

type U struct {
	id        uuid.UUID
	passId    string
	email     string
	userType  UE
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
