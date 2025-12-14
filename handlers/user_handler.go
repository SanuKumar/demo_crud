package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"demo_crud/config"
	"demo_crud/models"

	"github.com/gorilla/mux"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a user with name, email, age
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 200 {object} models.User
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result, err := config.DB.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	config.DB.QueryRow("SELECT id, name, email, age, created_at, updated_at FROM users WHERE id=?",
		id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)

	json.NewEncoder(w).Encode(user)
}

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve all users
// @Tags Users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query("SELECT id, name, email, age, created_at, updated_at FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user models.User
	err := config.DB.QueryRow("SELECT id, name, email, age, created_at, updated_at FROM users WHERE id=?",
		id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update user details by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated User Data"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	_, err := config.DB.Exec("UPDATE users SET name=?, email=?, age=? WHERE id=?",
		user.Name, user.Email, user.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	config.DB.QueryRow("SELECT id, name, email, age, created_at, updated_at FROM users WHERE id=?",
		id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.CreatedAt, &user.UpdatedAt)

	json.NewEncoder(w).Encode(user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete user by ID
// @Tags Users
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := config.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted successfully"))
}
