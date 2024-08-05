package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:8443", config)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	for {
		fmt.Print("Enter message: ")
		var input string
		fmt.Scanln(&input)

		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending to server:", err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Println("Received from server:", string(buffer[:n]))
	}
}
