package dns

import (
	"encoding/binary"
)

type Flags struct {
	QR     bool  // indicates whether the message is a query (0) or a reply (1)
	Opcode uint8 // 4 bits
	AA     bool  // Authoritative Answer
	TC     bool  // Truncation
	RD     bool  // Recursion Desired
	RA     bool  // Recursion Available
	Z      uint8 // 3 bits reserved for future use by DNSSEC
	Rcode  uint8 // 4 bits response code
}

func (*Flags) Marshal() []byte {
	marshaled := make([]byte, 2)

	marshaled[0] = 1 << 7 // QR

	return marshaled
}

type Header struct {
	ID    uint16 // transaction ID
	Flags *Flags // flags

	QDCount uint16 // number of entries in the question section
	AnCount uint16 // number of records in the answer section
	NsCount uint16 // number of records in the authority section
	ArCount uint16 // number of records in the additional section

}

func (m *Header) Marshal() []byte {
	marshaled := make([]byte, 12)
	binary.BigEndian.PutUint16(marshaled[0:2], m.ID)

	copy(marshaled[2:4], m.Flags.Marshal())

	binary.BigEndian.PutUint16(marshaled[4:6], m.QDCount)

	return marshaled
}
