package usecase

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

type IFaceUsecase interface {
	// auth
	Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error)
	Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error)

	// user
	GetListUser(ctx context.Context, params *request.ListUserQuery) ([]*response.ListUser, error)

	// nurse
	CreateNurse(ctx context.Context, req *request.CreateNurse) (*response.CreateNurse, error)

	// image
	UploadImage(ctx context.Context, file *multipart.FileHeader) (*response.ImageResponse, error)
}
