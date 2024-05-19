package services

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	Conn   *websocket.Conn
	RoomID uint
	UserID uint
}

var clients = make(map[*Client]bool)
var broadcast = make(chan Message)
var mutex = sync.Mutex{}

type Message struct {
	RoomID uint
	Data   string
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		return
	}
	defer conn.Close()

	client := &Client{Conn: conn}
	clients[client] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error reading message: %v", err)
			delete(clients, client)
			break
		}

		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			if client.RoomID == msg.RoomID {
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("error writing message: %v", err)
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
		mutex.Unlock()
	}
}
