package response

type ListPatient struct {
	IdentityNumber int    `json:"identityNumber" db:"identity_number"`
	Phone          string `json:"phoneNumber" db:"phone"`
	Name           string `json:"name" db:"name"`
	BirthDate      string `json:"birthDate" db:"birth_date"`
	Gender         string `json:"gender" db:"gender"`
	CreatedAt      string `json:"createdAt" db:"created_at"`
}
