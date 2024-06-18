package godb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx = context.Background()

var dbInstance *pgxpool.Pool

func InitDB() error {
	pool, err := pgxpool.NewWithConfig(ctx, Config())
	if err != nil {
		log.Fatal("unable to connect to database: ", err)
		return err
	}
	dbInstance = pool

	return nil
}

func GetDB() *pgxpool.Pool {
	return dbInstance
}

func CloseDB() {
	if dbInstance != nil {
		dbInstance.Close()
	}
}
