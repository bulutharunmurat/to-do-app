package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

const (
	// Initialize connection constants.
	HOST     = "localhost"
	PORT     = 5432
	DATABASE = "postgre"
	USER     = "postgre"
	PASSWORD = "pass"
)

func main() {
	// Capture connection properties.
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)
	fmt.Println(connectionString)
	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS to_do;")
	fmt.Println("Finished dropping table (if existed)")
}
