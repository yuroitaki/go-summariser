package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type BodyText struct {
	Text string
}

func main() {
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
	err = json.Unmarshal(body, &textBody)
	quit(err)
	log.Println(textBody.Text)
}

func quit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}