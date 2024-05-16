package repository

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
)

func (r *Repository) CreatePatient(ctx context.Context, data *model.Patient) (*model.Patient, error) {
	query := `INSERT INTO patients(identity_number, phone, name, birth_date, gender, identity_card_scan_img)
        VALUES (@identity_number, @phone, @name, @birth_date, @gender, @identity_card_scan_img) RETURNING patient_id, created_at`
	args := pgx.NamedArgs{
		"identity_number":        data.IdentityNumber,
		"phone":                  data.Phone,
		"name":                   data.Name,
		"birth_date":             data.BirthDate,
		"gender":                 data.Gender,
		"identity_card_scan_img": data.IdentityCardScanImg,
	}

	dest := []interface{}{
		&data.PatientID,
		&data.CreatedAt,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		if IsDuplicateError(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.HTTPStatusText(http.StatusConflict),
			})
		}
		return nil, err
	}

	return data, nil
}
