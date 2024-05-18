package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
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

func (r *Repository) FindPatients(ctx context.Context, params *request.ListPatientQuery) ([]*response.ListPatient, error) {
	query := `SELECT identity_number::bigint as identity_number, phone, name, to_char(birth_date, 'YYYY-MM-DD HH24:MI:SS.US') as birth_date, gender, to_char(created_at, 'YYYY-MM-DD HH24:MI:SS.US') as created_at FROM patients`

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	if params.IdentityNumber != nil {
		clause = append(clause, fmt.Sprintf(" identity_number = $%d", counter))
		args = append(args, *params.IdentityNumber)
		counter++
	} else {
		if params.Phone != nil {
			clause = append(clause, fmt.Sprintf(" phone like $%d", counter))
			args = append(args, "+"+*params.Phone+"%")
			counter++
		}
		if params.Name != nil {
			clause = append(clause, fmt.Sprintf(" name ilike $%d", counter))
			args = append(args, "%"+*params.Name+"%")
			counter++
		}
	}

	if counter > 1 {
		query += " WHERE" + strings.Join(clause, " AND")
	}

	orderClause := " ORDER BY created_at at time zone 'Asia/Jakarta' DESC"
	if params.CreatedAt != nil {
		if *params.CreatedAt == "asc" {
			orderClause = " ORDER BY created_at at time zone 'Asia/Jakarta' ASC"
		}
	}
	query += orderClause

	if params.Limit != nil && *params.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, params.Limit)
		counter++
	} else {
		query += fmt.Sprintf(" LIMIT $%d", counter)
		args = append(args, 5)
		counter++
	}

	if params.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, params.Offset)
		counter++
	} else {
		query += fmt.Sprintf(" OFFSET $%d", counter)
		args = append(args, 0)
		counter++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]*response.ListPatient, 0)
	for rows.Next() {
		p := &response.ListPatient{}
		err = rows.Scan(
			&p.IdentityNumber,
			&p.Phone,
			&p.Name,
			&p.BirthDate,
			&p.Gender,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}
