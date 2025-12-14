package main

import (
	"log"
	"net/http"

	"demo_crud/database"
	"demo_crud/routes"

	_ "demo_crud/docs" // swagger docs

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Go CRUD API
// @version 1.0
// @description CRUD API using Go and MySQL
// @host localhost:8080
// @BasePath /

func main() {
	database.ConnectDB()

	router := mux.NewRouter()

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	routes.RegisterRoutes(router)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
