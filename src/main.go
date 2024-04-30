package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	DBDriver = "mysql"
	DBPort   = "3306"
	DBName   = "workshop"
)

var DBUser = os.Getenv("MYSQL_DB_USER")
var DBHost = os.Getenv("MYSQL_DB_HOST")
var DBPass = os.Getenv("MYSQL_DB_PASSWORD")
var DBDatasource = DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName

func main() {
	// Prepare the database
	prepareDbHandler()

	// Create a new router
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/userslist", listUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
	r.HandleFunc("/user", createUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
