package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

// Variable definition of Database Connection
var db *pgx.Conn

// Database connection function

func connectDB() {
	var err error

	// Database Contact Information
	connStr := "postgres://admin:password@localhost:5432/go_web_db"

	//Connect to database
	db, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:❌", err)
	}

	fmt.Println("Database connection successful! ✅")
}

// CreateTables
func CreateTables() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price DECIMAL(10,2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE credit_cards (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		card_number TEXT NOT NULL,
		expiry_date TEXT NOT NULL,
		cvv TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE purchases (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		product_id INT REFERENCES products(id) ON DELETE CASCADE,
		purchase_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	`

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("❌ Failed to create tables:", err)
	}
	fmt.Println("✅ Tables created successfully!")
}