package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/dns-server-starter-go/dns"
)

// Ensures gofmt doesn't remove the "net" import in stage 1 (feel free to remove this!)
var _ = net.ListenUDP

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buff := make([]byte, 512)

	for {
		size, client, err := udpConn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		message := dns.Unmarshal(buff[:size])

		_, err = udpConn.WriteToUDP(message.Marshal(), client)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
