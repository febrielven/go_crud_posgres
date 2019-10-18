package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetPostgersDB() (*sql.DB, error) {
	// host := os.Getenv("POSTGRES_HOST")
	// user := os.Getenv("POSTGRES_USER")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// databaseName := os.Getenv("POSTGRES_DB_NAME")

	host := "localhost"
	user := "root"
	password := "123456"
	databaseName := "go-crud"

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)

	db, err := CreateConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
