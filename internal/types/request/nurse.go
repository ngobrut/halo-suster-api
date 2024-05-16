package request

import (
	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/constant"
)

type CreateNurse struct {
	NIP                 int               `json:"nip" validate:"required,nipNurse"`
	Name                string            `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImg string            `json:"identityCardScanImg" validate:"required,validUrl"`
	UserRole            constant.UserRole `json:"-"`
}

type UpdateNurse struct {
	NIP    int               `json:"nip" validate:"required,nipNurse"`
	Name   string            `json:"name" validate:"required,min=5,max=50"`
	UserID uuid.UUID         `json:"-"`
	Role   constant.UserRole `json:"-"`
}

type GrantNurseAccess struct {
	Password string            `json:"password" validate:"required,min=5,max=33"`
	UserID   uuid.UUID         `json:"-"`
	Role     constant.UserRole `json:"-"`
}
