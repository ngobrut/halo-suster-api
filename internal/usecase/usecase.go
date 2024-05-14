package usecase

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngobrut/halo-sus-api/config"
	"github.com/ngobrut/halo-sus-api/internal/repository"
)

type Usecase struct {
	cnf  config.Config
	db   *pgxpool.Pool
	repo repository.IFaceRepository
}

func New(cnf config.Config, db *pgxpool.Pool, repo repository.IFaceRepository) IFaceUsecase {
	return &Usecase{
		cnf:  cnf,
		db:   db,
		repo: repo,
	}
}
