package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	room string
}

var clients = make(map[string]*Client) //  koneksi klien

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	// Buat client baru
	client := &Client{conn: ws}

	// Tambahkan client ke map
	clients[client.conn.RemoteAddr().String()] = client

	// Loop untuk menerima pesan
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, client.conn.RemoteAddr().String())
			break
		}

		// Broadcast chat
		for _, c := range clients {
			if c.room == client.room {
				err := c.conn.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)
	http.ListenAndServe(":8081", nil)
}
