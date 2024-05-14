package repository

import (
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngobrut/halo-sus-api/config"
)

type Repository struct {
	cnf config.Config
	db  *pgxpool.Pool
}

func New(cnf config.Config, db *pgxpool.Pool) IFaceRepository {
	return &Repository{
		cnf: cnf,
		db:  db,
	}
}

func IsDuplicateError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key")
}

func IsRecordNotFound(err error) bool {
	return strings.Contains(err.Error(), "no rows in result set")
}
