package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/nrednav/cuid2"
	"github.com/rs/cors"
)

type PersonInfo struct {
	UserId   string
	Username string
}

var users map[string]string

func main() {
	fmt.Println("hello SEKAI!!")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	// initialize the global users map
	users = make(map[string]string)

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

func Register(writer http.ResponseWriter, request *http.Request) {
	requestUrl := request.URL
	params, err := url.ParseQuery(requestUrl.RawQuery)
	if err != nil {
		fmt.Printf("Error parsing URL parameters: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "{\"error\": \"%s\"}", err)
		return
	}

	usernameArr, ok := params["username"]
	if !ok {
		fmt.Println("username GET param not found")
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "{\"error\": \"username not found\"}")
		return
	}
	username := usernameArr[0]

	uid := cuid2.Generate()

	// Store in the users map
	users[uid] = username

	// The result json
	result := PersonInfo{
		UserId:   uid,
		Username: username,
	}

	writer.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error in JSON marshal: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "{\"error\": \"%s\"}", err)
		return
	}

	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "%s", jsonData)
}

func ValidateId(writer http.ResponseWriter, request *http.Request) {

}
