package request

import "github.com/ngobrut/halo-suster-api/constant"

type CreatePatient struct {
	IdentityNumber      int             `json:"identityNumber" validate:"required,idNum"`
	Phone               string          `json:"phoneNumber" validate:"required,min=10,max=15,phone"`
	Name                string          `json:"name" validate:"required,min=3,max=30"`
	BirthDate           string          `json:"birthDate" validate:"required,dateFormat"`
	Gender              constant.Gender `json:"gender" validate:"required,gender"`
	IdentityCardScanImg string          `json:"identityCardScanImg" validate:"required,validUrl"`
}
