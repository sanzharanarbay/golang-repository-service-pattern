package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbType := os.Getenv("DB_TYPE")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string

	// open connection
	db, err := sql.Open(dbType, dbUri)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Database!")
	// return the connection
	return db
}
