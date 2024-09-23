package service

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func InitializeDatabase() error {
	var err error
	Database, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=1111 dbname=zephyr_db sslmode=disable")
	if err != nil {
		return err
	}
	err = Database.Ping()
	if err != nil {
		return err
	}
	return nil
}
