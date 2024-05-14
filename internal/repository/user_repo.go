package repository

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
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
		QueryRow(ctx, "SELECT * FROM users WHERE nip = $1", nip).
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
		QueryRow(ctx, "SELECT * FROM users WHERE user_id = $1", userID).
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
