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
 _, err := godb.InitDB() // can also be dbInstance, err := godb.InitDB()
 if err != nil {
  panic(err)
 }
 defer godb.Close()

 /* check database connection and functionality on init */
 dberr := godb.Exec("CREATE TABLE IF NOT EXISTS temp_table (first VARCHAR($1) PRIMARY KEY);", "1")
 if dberr != nil {
  panic(dberr)
 }

 rows, err := godb.Query("SELECT * FROM temp_table")
 if err != nil {
  panic(err)
 }
 rows.Close()

 dberr = godb.Exec("DROP TABLE temp_table;")
 if dberr != nil {
  panic(dberr)
 }
}

```
