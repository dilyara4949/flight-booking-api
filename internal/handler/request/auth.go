package request

import (
	"github.com/google/uuid"
)

type Signup struct {
	Email    string    `json:"email"`
	Password string    `json:"password"`
	RoleID   uuid.UUID `json:"role_id"`
}
