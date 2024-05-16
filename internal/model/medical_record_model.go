package model

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	ID          uuid.UUID `db:"id"`
	PatientID   uuid.UUID `db:"patient_id"`
	UserID      uuid.UUID `db:"user_id"`
	Symptoms    string    `db:"symptoms"`
	Medications string    `db:"medications"`
	CreatedAt   time.Time `db:"created_at"`
}
