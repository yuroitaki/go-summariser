package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type BodyText struct {
	Text string
}

func main() {
	godotenv.Load()
	router := mux.NewRouter()
	listeningPort := "55100"
	listeningAddress := fmt.Sprintf("localhost:%s", listeningPort)

	router.HandleFunc("/summariser", summarise).Methods(http.MethodGet)
	log.Println("Summariser waiting for content...")
	http.ListenAndServe(listeningAddress, router)
}

func summarise(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	quit(err)

	var textBody BodyText
	// passing in the memory address of textBody as json Unmarshal accepts pointer
	err = json.Unmarshal(body, &textBody)
	quit(err)

	getOpenAi(textBody.Text)
}

func getOpenAi(text string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	text = text + "\ntl;dr:"
	
}

func quit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}