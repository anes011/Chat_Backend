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
	"golang.org/x/crypto/bcrypt"
)

type Body struct {
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
	body := Body{}
	json.NewDecoder(r.Body).Decode(&body)

	//Hashing the user password before saving it
	bytes, bError := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if bError != nil {
		http.Error(w, fmt.Sprintf("Error hashing password: %v", bError), 400)
		return
	}

	//Saving user into database
	err := controllers.CreateUser(&models.User{
		ID:        uuid.New(),
		UserName:  body.UserName,
		Email:     body.Email,
		Password:  string(bytes),
		Photo:     body.Photo,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), 400)
		return
	}

	utils.RespondWithJson(w, 201, &UserResponse{
		Success: true,
		Msg:     "User created successfully!",
	})
}
