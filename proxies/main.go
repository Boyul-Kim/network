package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("proxies")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	defer listener.Close()

	fmt.Println("TCP Server listening on port 8080")

	//continualy accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERR:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("CONN", conn.LocalAddr(), conn.RemoteAddr())
	//allocate buffer to read data into
	buffer := make([]byte, 4000)

	//read data sent from client
	for {
		n, err := conn.Read(buffer)
		// fmt.Println("N", n)
		if err != nil {
			fmt.Println("ERR:", err)
			return
		}
		// fmt.Println("BUFFER", buffer)
		fmt.Printf("Received: %s\n", buffer[:n])
	}
}
