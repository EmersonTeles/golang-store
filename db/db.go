package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)
func ConnectToPostgres() *sql.DB{
	connection := "host=localhost user=postgres dbname=alura_store password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}	