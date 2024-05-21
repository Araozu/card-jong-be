package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func LobbyWsConnect(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Print("read:", err)
			break
		}

		log.Printf("recv: %s, type: %d", message, mt)

		time.Sleep(10 * time.Second)

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Print("write:", err)
			break
		}
	}
}
