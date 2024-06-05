package controllers

import (
	"errors"

	"github.com/anes011/chat/pkg/database"
	"github.com/anes011/chat/pkg/database/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}

	db := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query(`SELECT id, user_name, email, pass, photo, created_at 
	FROM users WHERE email = $1`, email)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	found := false

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email,
			&user.Password, &user.Photo, &user.CreatedAt); err != nil {
			return nil, err
		}

		found = true
	}

	if !found {
		return nil, errors.New("no user found with this email")
	}

	return &user, nil
}
