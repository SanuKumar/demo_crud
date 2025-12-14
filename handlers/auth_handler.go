package handlers

import (
	"demo_crud/utils"
	"encoding/json"
	"net/http"
)

// Dummy login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse contains JWT token
type LoginResponse struct {
	Token string `json:"token"`
}

// Login godoc
// @Summary Login user
// @Description Generate JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	json.NewDecoder(r.Body).Decode(&loginReq)

	// In real app, validate username/password from DB
	if loginReq.Username != "admin" || loginReq.Password != "password" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(1) // userID = 1
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
