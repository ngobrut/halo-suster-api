package response

import "github.com/google/uuid"

type CreateNurse struct {
	UserID uuid.UUID `json:"userId"`
	NIP    int       `json:"nip"`
	Name   string    `json:"name"`
}
