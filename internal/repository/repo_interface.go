package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/internal/model"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User) error
	FindOneUserByNIP(ctx context.Context, nip string) (*model.User, error)
	FindOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
}
