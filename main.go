package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	listeningPort := "55100"
	listeningAddress := fmt.Sprintf("0.0.0.0:%s", listeningPort)

	router.HandleFunc("/summariser", Summarise).Methods(http.MethodGet)
	log.Println("Summariser waiting for content...")
	http.ListenAndServe(listeningAddress, router)
}

func Summarise(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello world!")
}