package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	listeningPort := "55100"
	listeningAddress := fmt.Sprintf("localhost:%s", listeningPort)

	router.HandleFunc("/summariser", summarise).Methods(http.MethodPost)
	log.Println("Summariser waiting for content...")
	http.ListenAndServe(listeningAddress, router)
}