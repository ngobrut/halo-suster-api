package response

import "github.com/google/uuid"

type AuthResponse struct {
	UserID      uuid.UUID `json:"userId"`
	NIP         string    `json:"nip"`
	Name        string    `json:"name"`
	AccessToken string    `json:"accessToken,omitempty"`
}
