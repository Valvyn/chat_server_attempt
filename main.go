package main

import (
	"fmt"
	"net"
)

const serverPort = ":8080"

func main() {
	serverConnection, err := net.Listen("tcp", serverPort)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Server has been successfully started on port", serverPort)

	for {
		connection, err := serverConnection.Accept()
		if err != nil {
			panic(err.Error())
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	connection.Write([]byte("Welcome to the Chatserver"))
	connection.Close()
}
