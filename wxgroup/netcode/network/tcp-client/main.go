// tcp-client/main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(conn, text)
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Failed to read response:", err)
			break
		}
		fmt.Print(response)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Failed to read line:", err)
	}
}
