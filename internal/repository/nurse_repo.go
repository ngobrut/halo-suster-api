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
