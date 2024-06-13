# Godb

postgres + pgx

## Env file

Next variables has to be set via .env file in Your project:

- POSTGRES_USER

- POSTGRES_PASSWORD

- POSTGRES_HOST

- POSTGRES_DB

## Init and usage database

```golang

import "github.com/myelophone/godb"

main() {
    if err := godb.InitDB(); err != nil {
        panic(err)
    }
    defer godb.CloseDB()
    dbStore := godb.GetDB()
    _, dberr := dbStore.Exec(ctx, "CREATE TABLE IF NOT EXISTS temp_table (first VARCHAR(1) PRIMARY KEY);")
    if dberr != nil {
        panic(dberr)
    }
    _, dberr = dbStore.Exec(ctx, "DROP TABLE temp_table;")
    if dberr != nil {
        panic(dberr)
    }
}

```
