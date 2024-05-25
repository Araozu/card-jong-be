package main

import (
	"card-jong-be/controller"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("hello SEKAI!!")
	mainRouter := mux.NewRouter()
	httpRouter := mainRouter.PathPrefix("/api").Subrouter()
	wsRouter := mainRouter.PathPrefix("/ws").Subrouter()

	// HTTP routes
	httpRouter.HandleFunc("/register", controller.RegisterUser)
	httpRouter.HandleFunc("/validate", controller.ValidateId)
	httpRouter.HandleFunc("/lobby/new", controller.CreateLobby).Methods("POST")

	// WS routes
	wsRouter.HandleFunc("/lobby/connect", controller.LobbyWsConnect)

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

	handler := cors.Handler(mainRouter)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
