package response

import "github.com/google/uuid"

type ListUser struct {
	UserID    uuid.UUID `json:"userId" db:"user_id"`
	NIP       int       `json:"nip" db:"nip"`
	Name      string    `json:"name" db:"name"`
	CreatedAt string    `json:"createdAt" db:"created_at"`
}
