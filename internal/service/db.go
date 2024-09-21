package service

import (
	"database/sql"
	"zephyr-api-mod/internal/models"

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

func GetUserByUsername(username string) (*models.User, error) {
	stmt, err := Database.Prepare("SELECT id, username, role, password, code FROM users WHERE username = $1")
	if err != nil {
		return nil, err
	}
	res := stmt.QueryRow(username)
	var user models.User
	err = res.Scan(&user.Id, &user.Username, &user.Role, &user.Password, &user.Code)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
