package redis

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/auth/internal/domain"
)

func makeSessionKey(s domain.Session) string {
	return makeIDKey(s.ID())
}

func makeIDKey(id uuid.UUID) string {
	return fmt.Sprintf("sessions:%s", id.String())
}

func makeTokenKey(token domain.SessionToken) string {
	return fmt.Sprintf("sessions:token:%s", token)
}
