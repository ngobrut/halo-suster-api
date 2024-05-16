package repository

import (
	"context"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
)

func (r *Repository) CreateMedicalRecord(ctx context.Context, req *request.CreateMedicalRecord) (*model.MedicalRecord, error) {
	query := `WITH patient_cte AS (
            SELECT patient_id  FROM patients p WHERE identity_number = @identity_number
        )
        INSERT INTO medical_records (patient_id, user_id, symptoms, medications)
        SELECT pc.patient_id, @user_id, @symptoms, @medications  FROM patient_cte pc
        RETURNING id, patient_id, user_id, symptoms, medications, created_at`
	args := pgx.NamedArgs{
		"identity_number": strconv.Itoa(req.IdentityNumber),
		"user_id":         req.UserID,
		"symptoms":        req.Symptoms,
		"medications":     req.Medications,
	}

	mr := &model.MedicalRecord{}

	err := r.db.QueryRow(ctx, query, args).Scan(&mr.ID, &mr.PatientID, &mr.UserID, &mr.Symptoms, &mr.Medications, &mr.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusBadRequest,
				Message:  "identityNumber is not exist",
			})
		}
		return nil, err
	}

	return mr, nil
}
