package controller

import (
	"github.com/nrednav/cuid2"
	"net/http"
)

var Users map[string]string = make(map[string]string)

func Register(username string) string {
	uid := cuid2.Generate()

	// Store in the users map
	Users[uid] = username

	return uid
}

func ValidateId(writer http.ResponseWriter, request *http.Request) {
	if AuthHeaderIsValid(request.Header.Get("Authorization")) {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}
