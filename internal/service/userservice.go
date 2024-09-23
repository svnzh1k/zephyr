package service

import "zephyr-api-mod/internal/models"

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
