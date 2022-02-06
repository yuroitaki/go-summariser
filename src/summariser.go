package main

import (
	// "context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/joho/godotenv"
	"database/sql"
	// gogpt "github.com/sashabaranov/go-gpt3"
)

type BodyText struct {
	Text string
	Temperature float32
	Engine string
	TopP float32
}

type JsonResponse struct {
	Summary string
	Text string
	Temperature float32
	Engine string
	TopP float32
}

const (
	MAX_WORD = 1000
	MAX_TOKEN = 500
)

func summarise(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	var textBody BodyText
	// passing in the memory address of textBody as json Unmarshal accepts pointer
	err = json.Unmarshal(body, &textBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if textBody.Text == "" {
		err = errors.New("invalid input body")
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	text := trimText(textBody.Text)

	validatedBody, err := validateTextBody(textBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	log.Println(text, validatedBody)

	db, err := setUpDB(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	summary, err := getSummary(text, validatedBody, db)
	if err == sql.ErrNoRows {
		log.Println("No existing summary found")
	} else if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	} else {
		log.Println("Summary:", summary)
	}

	// summarisedText, err := runGPT3(text, validatedBody)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(500)
	// 	return
	// }

	// response := JsonResponse{
	// 	Summary: summarisedText,
	// 	Temperature: validatedBody.Temperature,
	// 	TopP: validatedBody.TopP,
	// 	Engine: validatedBody.Engine,
	// 	Text: text,
	// }
	// json.NewEncoder(w).Encode(response)
}

func validateTextBody(textBody BodyText) (BodyText, error) {
	if textBody.Temperature < 0 || textBody.Temperature > 1 {
		return textBody, errors.New("temperature set cannot be higher than 1 or lower than 0")
	}
	if textBody.TopP < 0 || textBody.TopP > 1 {
		return textBody, errors.New("top_p set cannot be higher than 1 or lower than 0")
	}
	validEngines := map[string]bool {
		"text-davinci-001": true,
		"text-curie-001": true,
		"text-babbage-001": true,
		"text-ada-001": true,
	}
	if _, ok := validEngines[textBody.Engine]; !ok {
		return textBody, errors.New("engine submitted is not valid, must be either text-davinci-001, text-curie-001, text-babbage-001 or text-ada-001")
	}
	return textBody, nil
}

func trimText(text string) string {
	log.Printf("Trimming body text to be under %d words", MAX_WORD)
	var trimmedWords []string
	words := strings.Fields(text)
	if len(words) > MAX_WORD {
		trimmedWords = words[:MAX_WORD]	
	} else {
		trimmedWords = words
	}
	trimmedText := strings.Join(trimmedWords, " ")
	log.Printf("Trimmed text: %s", trimmedText)
	return trimmedText
}

// func runGPT3(text string, validatedBody BodyText) (string, error) {
// 	ctx := context.Background()
// 	apiKey := os.Getenv("OPENAI_API_KEY")
// 	text = text + "\ntl;dr:"
	
// 	client := gogpt.NewClient(apiKey)
// 	req := gogpt.CompletionRequest{
// 		Prompt: text,
// 		MaxTokens: MAX_TOKEN,
// 		Temperature: validatedBody.Temperature,
// 		TopP: validatedBody.TopP,
// 	}
// 	resp, err := client.CreateCompletion(ctx, validatedBody.Engine, req)
// 	if len(resp.Choices) == 0 {
// 		return "", errors.New("API response error")
// 	}
// 	return resp.Choices[0].Text, err
// }