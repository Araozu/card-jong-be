package main

import (
	"card-jong-be/controller"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type PersonInfo struct {
	UserId   string
	Username string
}

func main() {
	fmt.Println("hello SEKAI!!")
	mainRouter := mux.NewRouter()
	httpRouter := mainRouter.PathPrefix("/api").Subrouter()
	wsRouter := mainRouter.PathPrefix("/ws").Subrouter()

	// HTTP routes
	httpRouter.HandleFunc("/register", Register)
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

func Register(writer http.ResponseWriter, request *http.Request) {
	requestUrl := request.URL
	params, err := url.ParseQuery(requestUrl.RawQuery)
	if err != nil {
		controller.WriteError(err, "Error parsing URL parameters", &writer)
		return
	}

	usernameArr, ok := params["username"]
	if !ok {
		controller.WriteError(err, "username not found", &writer)
		return
	}
	username := usernameArr[0]

	// The result json
	result := PersonInfo{
		UserId:   controller.Register(username),
		Username: username,
	}

	writer.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(result)
	if err != nil {
		controller.WriteError(err, "Error serializing JSON", &writer)
		return
	}

	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "%s", jsonData)
}
