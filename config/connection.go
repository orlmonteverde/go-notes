package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// GetConnection return a connection to de database
func GetConnection() *sql.DB {
	dns := os.Getenv("DATABASE_URL") + "?sslmode=disable"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}
	return db
}
