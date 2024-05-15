package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/internal/model"
	"github.com/ngobrut/halo-suster-api/internal/types/request"
	"github.com/ngobrut/halo-suster-api/internal/types/response"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User) error
	FindOneUserByNIP(ctx context.Context, nip string) (*model.User, error)
	FindOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	FindUsers(ctx context.Context, params *request.ListUserQuery) ([]*response.ListUser, error)

	// nurse
	CreateNurse(ctx context.Context, data *model.User) error
	UpdateNurse(ctx context.Context, req *request.UpdateNurse) error
	DeleteNurse(ctx context.Context, userID uuid.UUID) error
	GrantNurseAccess(ctx context.Context, req *request.GrantNurseAccess) error
}
