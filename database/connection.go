package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	dbHost := "localhost"
	dbPort := "3336"
	dbUser := "root"
	dbPass := "1234"
	dbName := "usersdata"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println("dbURL=======> ", dbURI)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}
