package request

import "github.com/ngobrut/halo-suster-api/constant"

type CreateNurse struct {
	NIP                 int               `json:"nip" validate:"required,nipLen"`
	Name                string            `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImg string            `json:"identityCardScanImg" validate:"required,validUrl"`
	UserRole            constant.UserRole `json:"-"`
}