package response

type IdentityDetail struct {
	IdentityNumber      int    `json:"identityNumber" db:"identity_number"`
	Phone               string `json:"phoneNumber" db:"phone"`
	Name                string `json:"name" db:"name"`
	BirthDate           string `json:"birthDate" db:"birth_date"`
	Gender              string `json:"gender" db:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg" db:"identity_card_scan_img"`
}

type CreatedBy struct {
	NIP    string `json:"nip" db:"user_nip"`
	Name   string `json:"name" db:"user_name"`
	UserID string `json:"userId" db:"user_id"`
}

type ListMedicalRecord struct {
	IdentityDetail IdentityDetail `json:"identityDetail" db:"-"`
	Symptoms       string         `json:"symptoms" db:"symptoms"`
	Medications    string         `json:"medications" db:"medications"`
	CreatedAt      string         `json:"createdAt" db:"created_at"`
	CreatedBy      CreatedBy      `json:"createdBy" db:"-"`
}
