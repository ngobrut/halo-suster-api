package repository

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(ctx context.Context, data *model.User) error {
	query := `INSERT INTO users(name, nip, password, role) VALUES (@name, @nip, @password, @role) RETURNING user_id, name, nip`
	args := pgx.NamedArgs{
		"name":     data.Name,
		"nip":      data.NIP,
		"password": data.Password,
		"role":     data.Role,
	}

	if data.Role == constant.UserRoleNurse {
		query = `INSERT INTO users(name, nip, role) VALUES (@name, @nip, @password, @role) RETURNING user_id, name, nip`
		args = pgx.NamedArgs{
			"name": data.Name,
			"nip":  data.NIP,
			"role": data.Role,
		}
	}

	dest := []interface{}{
		&data.UserID,
		&data.Name,
		&data.NIP,
	}

	err := r.db.QueryRow(ctx, query, args).Scan(dest...)
	if err != nil {
		if IsDuplicateError(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.HTTPStatusText(http.StatusConflict),
			})
		}

		return err
	}

	return nil
}

// FindOneUserByNIP implements IFaceRepository.
func (r *Repository) FindOneUserByNIP(ctx context.Context, nip string) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow(ctx, "SELECT * FROM users WHERE nip = $1 and deleted_at IS NULL", nip).
		Scan(
			&res.UserID,
			&res.NIP,
			&res.Name,
			&res.Password,
			&res.Role,
			&res.IdentityCardScanImg,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, err
	}

	return res, nil
}

// FindOneUserByID implements IFaceRepository.
func (r *Repository) FindOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	res := &model.User{}

	err := r.db.
		QueryRow(ctx, "SELECT * FROM users WHERE user_id = $1 and deleted_at IS NULL", userID).
		Scan(
			&res.UserID,
			&res.NIP,
			&res.Name,
			&res.Password,
			&res.Role,
			&res.IdentityCardScanImg,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)

	if err != nil {
		if err == pgx.ErrNoRows {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusNotFound,
				Message:  constant.HTTPStatusText(http.StatusNotFound),
			})
		}

		return nil, err
	}

	return res, nil
}

// FindUsers implements IFaceRepository
func (r *Repository) FindUsers(ctx context.Context, params *request.ListUserQuery) ([]*response.ListUser, error) {
	query := `SELECT user_id, nip::bigint as nip, name, to_char(created_at, 'YYYY-MM-DD HH24:MI:SS') as created_at
		FROM users WHERE deleted_at IS NULL`

	var clause = make([]string, 0)
	var args = make([]interface{}, 0)
	var counter int = 1

	if params.UserID != nil {
		clause = append(clause, fmt.Sprintf(" user_id = $%d", counter))
		args = append(args, params.UserID)
		counter++
	} else {
		if params.Name != nil {
			clause = append(clause, fmt.Sprintf(" name ilike $%d", counter))
			args = append(args, "%"+*params.Name+"%")
			counter++
		}
		if params.NIP != nil {
			clause = append(clause, fmt.Sprintf(" nip like $%d", counter))
			args = append(args, *params.NIP+"%")
			counter++
		}
		if params.Role != nil {
			if *params.Role == "it" {
				clause = append(clause, fmt.Sprintf(" role = $%d", counter))
				args = append(args, "it")
				counter++
			}
			if *params.Role == "nurse" {
				clause = append(clause, fmt.Sprintf(" role = $%d", counter))
				args = append(args, "nurse")
				counter++
			}

		}
	}

	if counter > 1 {
		query += " AND" + strings.Join(clause, " AND")
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

	res := make([]*response.ListUser, 0)
	for rows.Next() {
		u := &response.ListUser{}
		err = rows.Scan(
			&u.UserID,
			&u.NIP,
			&u.Name,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, u)
	}

	return res, nil
}
