package godb

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx = context.Background()

var DBPOOL *pgxpool.Pool

func InitDB() {
	connPool, err := pgxpool.NewWithConfig(ctx, Config())
	if err != nil {
		log.Fatal("unable to connect to database: ", err)
	}

	err = connPool.Ping(ctx)
	if err != nil {
		panic(err)
	}

	DBPOOL = connPool

	connection, err := connPool.Acquire(ctx)
	if err != nil {
		log.Fatal("error while acquiring connection from the database pool!")
	}
	defer connection.Release()

	err = connection.Ping(ctx)
	if err != nil {
		log.Fatal("could not ping database")
	}

	//defer connPool.Close()
}

/*
-- Use the quote_literal function to escape user input
SELECT * FROM users WHERE username = quote_literal(input_username);
*/
