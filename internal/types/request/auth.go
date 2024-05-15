package request

import "github.com/ngobrut/halo-suster-api/constant"

type Register struct {
	NIP      int    `json:"nip" validate:"required,nipIt"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type Login struct {
	NIP      int               `json:"nip" validate:"required,nipLen"`
	Password string            `json:"password" validate:"required,min=5,max=33"`
	UserRole constant.UserRole `json:"-"`
}
