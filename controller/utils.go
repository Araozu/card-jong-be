package controller

import (
	"fmt"
	"net/http"
	"strings"
)

func WriteError(err error, message string, writer *http.ResponseWriter) {
	fmt.Printf("Error: %s\n", err)
	(*writer).WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(*writer, "{\"error\": \"%s\"}", message)
}

func AuthHeaderIsValid(authHeader string) bool {
	// (try to) get the Bearer token
	reqToken := authHeader
	if !strings.HasPrefix(reqToken, "Bearer ") {
		return false
	}

	bearerToken := reqToken[7:]

	// Check that the token is in the global map
	_, ok := Users[bearerToken]

	return ok
}
