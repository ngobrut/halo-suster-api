package usecase

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/ngobrut/halo-suster-api/constant"
	"github.com/ngobrut/halo-suster-api/internal/custom_error"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

func (u *Usecase) CreatePatient(ctx context.Context, req *request.CreatePatient) error {
	var birthDate time.Time

	for _, format := range constant.DateTime {
		if bd, err := time.Parse(format, req.BirthDate); err == nil {
			birthDate = bd
		}
	}

	patient := &model.Patient{
		IdentityNumber:      strconv.Itoa(req.IdentityNumber),
		Phone:               req.Phone,
		Name:                req.Name,
		BirthDate:           birthDate,
		Gender:              string(req.Gender),
		IdentityCardScanImg: req.IdentityCardScanImg,
	}

	_, err := u.repo.CreatePatient(ctx, patient)
	if err != nil {
		return err
	}

	return nil

}

func (u *Usecase) GetListPatient(ctx context.Context, params *request.ListPatientQuery) ([]*response.ListPatient, error) {
	res, err := u.repo.FindPatients(ctx, params)
	if err != nil {
		return nil, custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusInternalServerError,
			Message:  constant.HTTPStatusText(http.StatusInternalServerError),
		})
	}

	return res, nil
}
