package request

import "github.com/google/uuid"

type CreateMedicalRecord struct {
	IdentityNumber int       `json:"identityNumber" validate:"required,idNum"`
	Symptoms       string    `json:"symptoms" validate:"required,min=1,max=2000"`
	Medications    string    `json:"medications" validate:"required,min=1,max=2000"`
	UserID         uuid.UUID `json:"-"`
}

type ListMedicalRecordQuery struct {
	IdentityNumber *string
	UserID         *string
	NIP            *string
	Limit          *int
	Offset         *int
	CreatedAt      *string
}
