package repository

import (
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngobrut/halo-suster-api/config"
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
	return err == pgx.ErrNoRows
}
