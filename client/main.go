package main

import (
	"fmt"
	"net"
)

func main() {
	//connect
	conn, err := net.Dial("tcp", "localhost:8080")
	fmt.Println("CONN", conn.LocalAddr(), conn.RemoteAddr())
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}

	defer conn.Close()

	//send data to tcp server
	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
}
