package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func prepareDbHandler() error {
	db, err := sql.Open(DBDriver, DBDatasource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return CreateTable(db)
}

func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func ListUsers(db *sql.DB) ([]User, error) {
	// Execute the SQL query to fetch all users from the database
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to store the users
	users := []User{}

	// Iterate over the rows returned by the query
	for rows.Next() {
		var user User
		// Scan the values from the row into the user object
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		// Append the user to the slice
		users = append(users, user)
	}

	// Check for any errors during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// Return the users slice
	return users, nil
}

// CreateUser inserts a new user into the database with the given name and email.
func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

// GetUser retrieves a user from the database based on the given ID.
func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates the name and email of a user in the database based on the given ID.
func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database based on the given ID.
func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
