package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type UserRow struct {
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
	ID        uint
}

func DBinit() {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	rows, err := db.Query("Select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user UserRow
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", user)
	}

	defer db.Close()
}
