package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anes011/chat/pkg/controllers"
	"github.com/anes011/chat/pkg/database/models"
	"github.com/anes011/chat/pkg/utils"
	"github.com/google/uuid"
)

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type UserResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithJson(w, 400, &UserResponse{
			Success: false,
			Msg:     fmt.Sprintf("Error parsing json: %v", err),
		})
	}

	err := controllers.CreateUser(&models.User{
		ID:        uuid.New(),
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		Photo:     user.Photo,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		utils.RespondWithJson(w, 400, err)
	} else {
		utils.RespondWithJson(w, 201, &UserResponse{
			Success: true,
			Msg:     "User created successfully!",
		})
	}
}
