// udp-server/main.go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		log.Fatal("Failed to resolve address:", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Failed to receive message:", err)
			continue
		}
		message := buf[:n]
		log.Printf("Received from %s: %s", remoteAddr, message)

		_, err = conn.WriteToUDP([]byte(fmt.Sprintf("Echo: %s", message)), remoteAddr)
		if err != nil {
			log.Println("Failed to send message:", err)
		}
	}
}
