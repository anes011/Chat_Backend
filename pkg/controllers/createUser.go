package controllers

import (
	"github.com/anes011/chat/pkg/database"
	"github.com/anes011/chat/pkg/database/models"
)

func CreateUser(user *models.User) error {
	db := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query(`INSERT INTO users (id, user_name, email, pass, 
	photo, created_at) VALUES ($1, $2, $3, $4, $5, $6)`, user.ID,
		user.UserName, user.Email, user.Password, user.Photo, user.CreatedAt)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
