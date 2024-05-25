package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nrednav/cuid2"
)

var Users map[string]string = make(map[string]string)

type PersonInfo struct {
	UserId   string
	Username string
}

func Register(username string) string {
	uid := cuid2.Generate()

	// Store in the users map
	Users[uid] = username

	return uid
}

func RegisterUser(writer http.ResponseWriter, request *http.Request) {

	requestUrl := request.URL
	params, err := url.ParseQuery(requestUrl.RawQuery)
	if err != nil {
		WriteError(err, "Error parsing URL parameters", &writer)
		return
	}

	usernameArr, ok := params["username"]
	if !ok {
		WriteError(err, "username not found", &writer)
		return
	}
	username := usernameArr[0]

	// The result json
	result := PersonInfo{
		UserId:   Register(username),
		Username: username,
	}

	writer.Header().Set("Content-Type", "application/json")

	jsonData, err := json.Marshal(result)
	if err != nil {
		WriteError(err, "Error serializing JSON", &writer)
		return
	}

	writer.WriteHeader(http.StatusOK)

	fmt.Fprintf(writer, "%s", jsonData)
}

func ValidateId(writer http.ResponseWriter, request *http.Request) {
	if AuthHeaderIsValid(request.Header.Get("Authorization")) {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}
