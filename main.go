package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("hello SEKAI!!")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/register", Register)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Origin", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	handler := cors.Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

type PersonInfo struct {
	useruuid int
	username string
}

func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "{\"yes\":\"and\"}")
}
