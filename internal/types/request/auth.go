package request

import "github.com/ngobrut/halo-suster-api/constant"

type Register struct {
	NIP      string `json:"nip" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Login struct {
	NIP      string            `json:"nip" validate:"required"`
	Password string            `json:"password" validate:"required"`
	UserRole constant.UserRole `json:"-"`
}
