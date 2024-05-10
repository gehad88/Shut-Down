package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()
	fmt.Println("Server started, listening on port 9090")

	fmt.Print("Enter command to send to clients: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn, command)
	}
}

func handleConnection(conn net.Conn, command string) {
	defer conn.Close()
	fmt.Println("New client connected:", conn.RemoteAddr().String())
	fmt.Println("Sending command to client:", command)
	_, err := conn.Write([]byte(command))
	if err != nil {
		fmt.Println("Error sending command to client:", err.Error())
		return
	}
}
