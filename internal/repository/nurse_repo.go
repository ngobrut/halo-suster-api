package repository

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
)

// CreateNurse implements IFaceRepository.
func (r *Repository) CreateNurse(ctx context.Context, data *model.User) error {
	query := `INSERT INTO users(name, nip, identity_card_scan_img, role)
        VALUES (@name, @nip, @identity_card_scan_img, @role) RETURNING user_id, name, nip, identity_card_scan_img`
	args := pgx.NamedArgs{
		"name":                   data.Name,
		"nip":                    data.NIP,
		"identity_card_scan_img": data.IdentityCardScanImg,
		"role":                   data.Role,
	}

	dest := []interface{}{
		&data.UserID,
		&data.Name,
		&data.NIP,
		&data.IdentityCardScanImg,
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

func (r *Repository) UpdateNurse(ctx context.Context, data *model.User) error {
	query := `UPDATE users
		SET nip = @nip, name = @name, updated_at = @updated_at
		WHERE user_id = @user_id`
	args := pgx.NamedArgs{
		"nip":        data.NIP,
		"name":       data.Name,
		"updated_at": data.UpdatedAt,
		"user_id":    data.UserID,
	}

	_, err := r.db.Exec(ctx, query, args)
	if err != nil {
		if IsDuplicateError(err) {
			return custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Message:  constant.HTTPStatusText(http.StatusConflict),
			})
		}
		return custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return nil
}
