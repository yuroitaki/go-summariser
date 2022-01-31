package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func setUpDB(user string, password string, dbName string) (*sql.DB, error) {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", conn)
	return db, err
}

