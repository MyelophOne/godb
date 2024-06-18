package godb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx = context.Background()

func InitDB() (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(ctx, Config())
	if err != nil {
		log.Fatal("unable to connect to database: ", err)
		return pool, err
	}

	return pool, nil
}
