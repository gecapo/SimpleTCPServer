package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host     = "localhost"
	port     = "6969"
	connType = "tcp"
)

func main() {
	l, err := net.Listen(connType, host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + host + ":" + port)

	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error when accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error when reading from connection:", err.Error())
	}

	conn.Write([]byte("Message received.\n"))

	conn.Close()
}
