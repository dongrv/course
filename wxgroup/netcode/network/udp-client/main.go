// udp-client/main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			log.Println("Failed to receive message:", err)
			break
		}
		fmt.Println(string(response[:n]))
	}
	if err := scanner.Err(); err != nil {
		log.Println("Failed to read line:", err)
	}
}
