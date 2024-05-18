package repository

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
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
				HTTPCode: http.StatusNotFound,
				Message:  "identityNumber is not exist",
			})
		}
		return nil, err
	}

	return mr, nil
}

func (r *Repository) FindMedicalRecords(ctx context.Context, params *request.ListMedicalRecordQuery) ([]*response.ListMedicalRecord, error) {
	query := `SELECT
			p.identity_number::bigint as identity_number,
			p.phone,
			p.name,
			to_char(p.birth_date, 'YYYY-MM-DD HH24:MI:SS.US') as birth_date,
			p.gender,
			p.identity_card_scan_img,
			mr.symptoms,
			mr.medications,
			to_char(mr.created_at, 'YYYY-MM-DD HH24:MI:SS.US') as created_at,
			u.nip::bigint as user_nip,
			u.name as user_name,
			u.user_id
		FROM medical_records mr
		LEFT JOIN patients p ON mr.patient_id = p.patient_id
		LEFT JOIN users u ON mr.user_id = u.user_id
		`
	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	if params.IdentityNumber != nil {
		clause = append(clause, fmt.Sprintf(" p.identity_number = $%d", counter))
		args = append(args, *params.IdentityNumber)
		counter++
	}
	if params.UserID != nil {
		clause = append(clause, fmt.Sprintf(" u.user_id = $%d", counter))
		args = append(args, *params.UserID)
		counter++
	}
	if params.NIP != nil {
		clause = append(clause, fmt.Sprintf(" u.nip = $%d", counter))
		args = append(args, *params.NIP)
		counter++
	}

	if counter > 1 {
		query += " WHERE" + strings.Join(clause, " AND")
	}

	orderClause := " ORDER BY mr.created_at at time zone 'Asia/Jakarta' DESC"
	if params.CreatedAt != nil {
		if *params.CreatedAt == "asc" {
			orderClause = " ORDER BY mr.created_at at time zone 'Asia/Jakarta' ASC"
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

	res := make([]*response.ListMedicalRecord, 0)
	for rows.Next() {
		mr := &response.ListMedicalRecord{}
		err = rows.Scan(
			&mr.IdentityDetail.IdentityNumber,
			&mr.IdentityDetail.Phone,
			&mr.IdentityDetail.Name,
			&mr.IdentityDetail.BirthDate,
			&mr.IdentityDetail.Gender,
			&mr.IdentityDetail.IdentityCardScanImg,
			&mr.Symptoms,
			&mr.Medications,
			&mr.CreatedAt,
			&mr.CreatedBy.NIP,
			&mr.CreatedBy.Name,
			&mr.CreatedBy.UserID,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, mr)
	}

	return res, nil
}
