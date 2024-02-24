package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    db, err := sql.Open("postgres", "postgres://username:password@localhost/database_name?sslmode=disable")
    if err != nil {
        return nil, err
    }
    return db, nil
}
