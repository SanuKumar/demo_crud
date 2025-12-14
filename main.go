package main

import (
	"log"
	"net/http"
	"os"

	"demo_crud/config"
	"demo_crud/routes"

	_ "demo_crud/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go CRUD API
// @version 1.0
// @description CRUD API using Go and MySQL
// @host localhost:8080
// @BasePath /

func main() {
	config.ConnectDB()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // default
	}

	router := mux.NewRouter()

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	routes.RegisterRoutes(router)

	log.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
