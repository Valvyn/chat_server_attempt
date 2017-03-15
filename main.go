package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const serverPort = ":8080"

var connectionList map[*websocket.Conn]int8

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", newWebSocket)
	http.ListenAndServe(serverPort, nil)
	fmt.Println("Server has been successfully started on port", serverPort)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	connectionList[conn] = 1

	err = handleMessages(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMessages(conn *websocket.Conn) error {
	for {
		fmt.Println(connectionList)
		for client := range connectionList {
			messageType, r, err := conn.NextReader()
			if err != nil {
				return err
			}
			w, err := client.NextWriter(messageType)
			if err != nil {
				return err
			}
			if _, err := io.Copy(w, r); err != nil {
				return err
			}
			if err := w.Close(); err != nil {
				return err
			}
		}
	}
}
