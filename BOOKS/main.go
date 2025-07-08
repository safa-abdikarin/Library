package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)



func main() {
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

	//INSERT DATA
	_, err = db.Exec("INSERT INTO books (title, author, year) VALUES ($1, $2, $3)",
	"GoLang 101", "Oreva", 2025)
	if err != nil {
		panic(err)
	}

	fmt.Println("Book added to the database successfully!")

	//LIST OF BOOKS
	rows, err := db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var author string
		var year int

		if err := rows.Scan(&id, &title, &author, &year); err != nil{
			panic(err)
		}
		fmt.Printf("ID: %d, TITLE:%s, AUTHOR:%s, YEAR: %d/n", id, title, author, year)
	}

	// UPDATE A BOOKS TITLE
	_, err = db.Exec("UPDATE books SET title = $1 WHERE title = $2", "GoLang 101 - Latest Edition", "GoLang 101")
	if err!= nil {
		panic(err)
	}
	fmt.Println("Title Updated Successfully!")

	//DELETE A BOOK BY ID
	_,err = db.Exec("DELETE FROM books WHERE id = $1", 1)
	if err!= nil {
		panic(err)
	}
	fmt.Println("User deleted successfully!") 

	
}