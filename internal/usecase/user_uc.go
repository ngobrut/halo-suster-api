package usecase

import (
	"context"
	"net/http"

	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

func (u *Usecase) GetListUser(ctx context.Context, params *request.ListUserQuery) ([]*response.ListUser, error) {
	res, err := u.repo.FindUsers(ctx, params)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}
	return res, nil
}
