package service

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID
	Restaurant uuid.UUID
	Products   []Product
	Customer   struct {
		ID      uuid.UUID
		Address string
	}
	CreatedAt time.Time

	Acceptance *struct {
		Manager    uuid.UUID
		AcceptedAt time.Time
	}

	Courier *struct {
		ID      uuid.UUID
		TakenAt time.Time
	}

	Payment *struct {
		Transaction string
		PayedAt     time.Time
	}

	Cancellation *struct {
		Canceller   uuid.UUID
		CancelledAt time.Time
	}

	CookingStartedAt  *time.Time
	DeliveryStartedAt *time.Time
	DeliveryCompleted *time.Time
}

type Product struct {
	ID          uuid.UUID
	Original    uuid.UUID
	Restaurant  uuid.UUID
	Name        string
	Picture     string
	Price       float64
	Description string
}
