package main

import (
	"fmt"
	"net"
	"os"
)

const (
	ConnectionPort = "6379"
	ConnectionType = "tcp"
	ConnectionHost = "localhost"
)

func main() {

	// Start a TCP connection on the specified port
	listener, err := net.Listen(ConnectionType, ConnectionHost+":"+ConnectionPort)
	if err != nil {
		fmt.Printf("error opening connection for %s on port: %s, err: %s", ConnectionType, ConnectionPort, err)
		os.Exit(1)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Printf("error in closing connection: %d", err)
		}
	}(listener)

	for {
		// Listen each connection request
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting: %s", err)
			os.Exit(1)
		}

		go parseAndHandleConnection(conn)
	}
}

func parseAndHandleConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	receivedBytes, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("error reading data: %s", err)
		return
	}
	fmt.Println(receivedBytes)
	fmt.Println(string(buf))
}
