package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/constant"
)

type User struct {
	UserID              uuid.UUID         `json:"user_id" db:"user_id"`
	NIP                 string            `json:"nip" db:"nip"`
	Name                string            `json:"name" db:"name"`
	Password            sql.NullString    `json:"-" db:"password"`
	Role                constant.UserRole `json:"role" db:"role"`
	IdentityCardScanImg sql.NullString    `json:"identity_card_scan_img" db:"identity_card_scan_img"`
	CreatedAt           time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at" db:"updated_at"`
	DeletedAt           *time.Time        `json:"-" db:"deleted_at"`
}
