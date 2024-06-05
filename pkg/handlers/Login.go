package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anes011/chat/pkg/controllers"
	"github.com/anes011/chat/pkg/database/models"
	"github.com/anes011/chat/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	Success bool `json:"success"`
	User    *models.User
}

func Login(w http.ResponseWriter, r *http.Request) {
	userCredentials := &UserCredentials{}
	json.NewDecoder(r.Body).Decode(&userCredentials)

	user, err := controllers.GetUserByEmail(userCredentials.Email)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), 400)
		return
	}

	value := CheckPasswordHash(userCredentials.Password, user.Password)

	if value {
		utils.RespondWithJson(w, 200, &LoginSuccessResponse{
			Success: true,
			User:    user,
		})
		return
	}

	http.Error(w, "Password is incorrect", 400)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
