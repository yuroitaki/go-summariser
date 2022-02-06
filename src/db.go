package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"crypto/sha256"
)

type Summarised struct {
	SummaryId string
	Summary string
	OriginalText string
	Temperature float32
	TopP float32
	Engine string
}

func setUpDB(user string, password string, dbName string) (*sql.DB, error) {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", conn)
	return db, err
}

func getSummary(text string, validatedBody BodyText, db *sql.DB) (string, error) {
	summaryIdString := fmt.Sprintf("%s%f%f%s", text, validatedBody.Temperature, validatedBody.TopP, validatedBody.Engine)
	summaryId := sha256hex(summaryIdString)

	var summary string
	sqlStatement := "SELECT summary FROM summary WHERE summary_id=$1"
	row := db.QueryRow(sqlStatement, summaryId)
	err := row.Scan(&summary)

	return summary, err
}

func sha256hex(payload string) string {
	payloadByte := []byte(payload)
	return fmt.Sprintf("%x", sha256.Sum256(payloadByte))
}