package controller

import (
	"encoding/json"
	"fmt"
	"github.com/nrednav/cuid2"
	"net/http"
)

// TODO: This struct should have a creation time
type Lobby struct {
	LobbyOwner   string
	LobbyPlayers [3]string
}

type LobbyResult struct {
	LobbyId string
}

// TODO: We should remove entries from this map when they expire.
// TODO: Define how long lobbies last
var lobbies = make(map[string]Lobby)

func CreateLobby(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	authOk := AuthHeaderIsValid(request.Header.Get("Authorization"))
	if !authOk {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	lobbyId := cuid2.Generate()

	result := LobbyResult{LobbyId: lobbyId}

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error in JSON marshal: %s\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "{\"error\": \"%s\"}", err)
		return
	}

	fmt.Fprintf(writer, "%s", jsonData)
}
