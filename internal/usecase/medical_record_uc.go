package usecase

import (
	"context"

	"github.com/ngobrut/halo-suster-api/internal/types/request"
)

func (u *Usecase) CreateMedicalRecord(ctx context.Context, req *request.CreateMedicalRecord) error {
	_, err := u.repo.CreateMedicalRecord(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
