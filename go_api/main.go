package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var userModel *UserModel

func main() {
	var err error
	// PostgreSQL connection string
	connStr := "postgres://postgres:postgres@localhost:5432/benchmark_db?sslmode=disable"

	// Connect to PostgreSQL
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		age INTEGER,
		password TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	userModel = &UserModel{DB: db}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", getAllUsers)
	mux.HandleFunc("GET /user", authMiddleware(login))
	mux.HandleFunc("POST /user", authMiddleware(register))
	mux.HandleFunc("PUT /user", authMiddleware(updateUser))
	mux.HandleFunc("DELETE /user", authMiddleware(deleteUser))

	http.ListenAndServe("localhost:3000", mux)
}
