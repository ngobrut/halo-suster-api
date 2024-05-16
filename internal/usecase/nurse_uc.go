package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
	"github.com/ngobrut/halo-suster-api/util"
)

func (u *Usecase) CreateNurse(ctx context.Context, req *request.CreateNurse) (*response.CreateNurse, error) {
	nurse := &model.User{
		NIP:                 strconv.Itoa(req.NIP),
		Name:                req.Name,
		IdentityCardScanImg: sql.NullString{String: req.IdentityCardScanImg, Valid: true},
		Role:                constant.UserRoleNurse,
	}

	err := u.repo.CreateNurse(ctx, nurse)
	if err != nil {
		return nil, err
	}

	nip, err := strconv.Atoi(nurse.NIP)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	res := &response.CreateNurse{
		UserID: nurse.UserID,
		Name:   nurse.Name,
		NIP:    nip,
	}
	return res, nil
}

func (u *Usecase) UpdateNurse(ctx context.Context, req *request.UpdateNurse) error {
	err := u.repo.UpdateNurse(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) DeleteNurse(ctx context.Context, userID uuid.UUID) error {
	err := u.repo.DeleteNurse(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GrantNurseAccess(ctx context.Context, req *request.GrantNurseAccess) error {
	pwd, err := util.HashPwd(u.cnf.BcryptSalt, []byte(req.Password))
	if err != nil {
		return err
	}
	req.Password = pwd

	err = u.repo.GrantNurseAccess(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
