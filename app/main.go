package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type DNSMessage struct {
	ID      uint16 // transaction ID
	QDCount uint16 // number of entries in the question section
	AnCount uint16 // number of records in the answer section
	NsCount uint16 // number of records in the authority section
	ArCount uint16 // number of records in the additional section

	Opcode uint8 // 4 bits
	Z      uint8 // 3 bits reserved for future use by DNSSEC
	Rcode  uint8 // 4 bits response code

	QR bool // indicates whether the message is a query (0) or a reply (1)
	AA bool // Authoritative Answer
	TC bool // Truncation
	RD bool // Recursion Desired
	RA bool // Recursion Available
}

func (m *DNSMessage) Marshal() []byte {
	marshaled := make([]byte, 12)
	binary.BigEndian.PutUint16(marshaled[0:2], m.ID)
	marshaled[2] = 1 << 7 // QR

	return marshaled

}

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

	req := make([]byte, 512)

	for {
		_, client, err := udpConn.ReadFromUDP(req)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		fmt.Println("Received data", req)

		reply := DNSMessage{
			ID:      1234,
			QR:      true,
			Opcode:  0,
			AA:      false,
			TC:      false,
			RD:      true,
			RA:      false,
			Z:       0,
			Rcode:   0,
			QDCount: 0,
			AnCount: 0,
			NsCount: 0,
			ArCount: 0,
		}

		_, err = udpConn.WriteToUDP(reply.Marshal(), client)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
