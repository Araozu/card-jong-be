package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type LobbyMsg struct {
	Action string `json:action`
	Value  string `json:value`
}

func LobbyWsConnect(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Print("read error:", err)
			break
		}

		log.Printf("recv: %s, type: %d", message, mt)

		var data LobbyMsg
		err = json.Unmarshal(message, &data)
		if err != nil {
			log.Print("json error:", err)
			break
		}

		switch data.Action {
		case "auth":
			err = validateUserId(mt, conn, data.Value)
		default:
			log.Print("no action :c")
		}

		if err != nil {
			log.Print("error:", err)
			break
		}
	}
}

func validateUserId(mt int, conn *websocket.Conn, userId string) error {
	_, ok := Users[userId]

	var responseJson LobbyMsg

	if ok {
		responseJson = LobbyMsg{Action: "auth", Value: "authorized"}
	} else {
		responseJson = LobbyMsg{Action: "auth", Value: "unauthorized"}
	}

	json, err := json.Marshal(responseJson)
	if err != nil {
		log.Print("json marshal: ", err)
		return err
	}

	err = conn.WriteMessage(mt, json)
	if err != nil {
		log.Print("write error: ", err)
	}
	return err
}
