package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ngobrut/halo-sus-api/config"
	"github.com/sirupsen/logrus"
)

type queryTracer struct {
	log *logrus.Logger
}

func (tracer *queryTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return ctx
}

func (tracer *queryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		tracer.log.Errorf("[error-query] %s\n", data)
	}
}

func NewDBClient(cnf config.Config, logger *logrus.Logger) (*pgxpool.Pool, error) {
	host := cnf.Postgres.Host
	port := cnf.Postgres.Port
	user := cnf.Postgres.User
	password := cnf.Postgres.Password
	dbname := cnf.Postgres.Database
	params := cnf.Postgres.Params
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, dbname, params)

	pgxconf, err := pgxpool.ParseConfig(uri)
	if err != nil {
		log.Printf("[error-parsing-db-config] %v\n", err)
		return nil, err
	}

	pgxconf.MaxConns = 20
	pgxconf.ConnConfig.Tracer = &queryTracer{
		log: logger,
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxconf)
	if err != nil {
		log.Printf("[error-db-connection] %v\n", err)
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Printf("[error-db-connection] %v\n", err)
		return nil, err
	}

	log.Println("[postgres-connected]")

	return pool, nil
}
