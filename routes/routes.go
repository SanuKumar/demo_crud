package routes

import (
	"demo_crud/handlers"
	"demo_crud/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Public route
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	protected.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	protected.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	protected.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	protected.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
}
