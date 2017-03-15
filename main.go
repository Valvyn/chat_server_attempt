package main

import (
	"fmt"
	"net/http"
)

const serverPort = ":8080"

func main() {
	fmt.Println("Server has been successfully started on port", serverPort)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", newWebSocket)
	http.ListenAndServe(serverPort, nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func newWebSocket(w http.ResponseWriter, r *http.Request) {
}
