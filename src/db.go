package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"crypto/sha256"
	"log"
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

func getSummary(summaryIdString string, db *sql.DB) (string, error) {
	summaryId := sha256hex(summaryIdString)
	log.Printf("Querying summary_id: %s from summary table",summaryId)

	var summary string
	sqlStatement := "SELECT summary FROM summary WHERE summary_id=$1"

	row := db.QueryRow(sqlStatement, summaryId)
	err := row.Scan(&summary)

	return summary, err
}

func insertSummary(text string, validatedBody BodyText, summarisedText string, summaryIdString string, db *sql.DB) error {
	summaryId := sha256hex(summaryIdString)
	log.Printf("Inserting summary_id: %s into summary table",summaryId)
	
	var lastId string
	sqlStatement := "INSERT INTO summary (summary_id, summary, original_text, temperature, top_p, engine) VALUES ($1, $2, $3, $4, $5, $6) RETURNING summary_id"
	err := db.QueryRow(sqlStatement, summaryId, summarisedText, text, validatedBody.Temperature, validatedBody.TopP, validatedBody.Engine).Scan(&lastId)
	return err
}

func sha256hex(payload string) string {
	payloadByte := []byte(payload)
	return fmt.Sprintf("%x", sha256.Sum256(payloadByte))
}