package usecase

import (
	"context"
	"mime/multipart"

	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

// UploadImage implements IFaceUsecase.
func (u *Usecase) UploadImage(ctx context.Context, file *multipart.FileHeader) (*response.ImageResponse, error) {
	res, err := u.aws.UploadFile(ctx, file)
	if err != nil {
		return nil, err
	}

	image := &response.ImageResponse{
		ImageURL: res.Location,
	}

	return image, nil
}
