package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const serverPort = ":8080"

var connections []*websocket.Conn
var incomingMessages chan io.Reader

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

	connections = append(connections, conn)

	go handleIncomingMessages(conn)
	go handleOutgoingMessages()
}

func handleOutgoingMessages() {
	for r := range incomingMessages {
		for _, client := range connections {
			w, err := client.NextWriter(1)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := io.Copy(w, r); err != nil {
				log.Fatal(err)
			}
			if err := w.Close(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
func handleIncomingMessages(conn *websocket.Conn) {
	for {
		_, r, err := conn.NextReader()
		if err != nil {
			log.Fatal(err)
		}
		incomingMessages <- r
	}
}
