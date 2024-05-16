package usecase

import (
	"context"
	"net/http"

	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

func (u *Usecase) CreateMedicalRecord(ctx context.Context, req *request.CreateMedicalRecord) error {
	_, err := u.repo.CreateMedicalRecord(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) GetListMedicalRecord(ctx context.Context, params *request.ListMedicalRecordQuery) ([]*response.ListMedicalRecord, error) {
	res, err := u.repo.FindMedicalRecords(ctx, params)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return res, nil
}
