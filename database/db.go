package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := "root:QazWsx@909@tcp(127.0.0.1:3306)/demo_crud?parseTime=true"

	DB, err = sql.Open("mysql", dsn)
	log.Println("**", err)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed")
	}
	log.Println("Database connected")
}
