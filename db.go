package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

var DB *pgxpool.Pool

func ConnectDB() {
	errr := godotenv.Load()
	if errr != nil {
		return
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/postgres"
	}
	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected to database")

}
