package godb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbInstance *pgxpool.Pool

var ctx = context.Background()

func InitDB() (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(ctx, Config())
	if err != nil {
		return nil, fmt.Errorf("error creating database connection pool: %w", err)
	}
	dbInstance = pool
	return pool, nil
}

// Query executes a query against the database pool.
func Query(query string, args ...interface{}) (pgx.Rows, error) {
	return dbInstance.Query(ctx, query, args...)
}

// Exec executes a statement (like INSERT, UPDATE, DELETE) against the database pool.
func Exec(query string, args ...interface{}) error {
	_, err := dbInstance.Exec(ctx, query, args...)
	return err
}

func Close() {
	dbInstance.Close()
}
