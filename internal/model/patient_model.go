package model

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	PatientID           uuid.UUID `json:"patient_id" db:"patient_id"`
	IdentityNumber      string    `json:"identityNumber" db:"identity_number"`
	Phone               string    `json:"phoneNumber" db:"phone"`
	Name                string    `json:"name" db:"name"`
	BirthDate           time.Time `json:"birthDate" db:"birth_date"`
	Gender              string    `json:"gender" db:"gender"`
	IdentityCardScanImg string    `json:"identityCardScanImg" db:"identity_card_scan_img"`
	CreatedAt           time.Time `json:"createdAt" db:"created_at"`
}
