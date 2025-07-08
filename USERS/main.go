package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)



func main() {
	username := "safa124"
    inputPassword := "password"
	
	//CONNECTION STRING
	connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "password", "Safa",
    )

	//open the database
		db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database successfully!")

	// ADD NEW USER - Insert Data
	_,err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", "safa124", "password")
	if err != nil {
		panic(err)
	}
	fmt.Println("User added successfully!")

	// VERIFY A USERS LOGIN CREDENTIALS
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		fmt.Println("User not found!")
	} else if storedPassword == inputPassword {
		 fmt.Println("Login Successful!")
	} else {
		fmt.Println("Incorrect Password")
	}



}