package controller

import (
	"strings"
)

func AuthHeaderIsValid(users *map[string]string, authHeader string) bool {
	// (try to) get the Bearer token
	reqToken := authHeader
	if !strings.HasPrefix(reqToken, "Bearer ") {
		return false
	}

	bearerToken := reqToken[7:]

	// Check that the token is in the global map
	_, ok := (*users)[bearerToken]

	return ok
}
