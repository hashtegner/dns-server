package dns

import (
	"encoding/binary"
)

type Answer struct {
	Name   *Name
	Type   uint16 // 2 bytes (A, NS, CNAME, etc.)
	Class  uint16 // 2 bytes (IN, CH, etc.)
	TTL    uint32 // 4 bytes
	Length uint16 // 2 bytes
	Data   []byte
}

func (a *Answer) Marshal() []byte {
	marshaled := make([]byte, 0)

	marshaled = append(marshaled, a.Name.Marshal()...)
	marshaled = binary.BigEndian.AppendUint16(marshaled, a.Type)
	marshaled = binary.BigEndian.AppendUint16(marshaled, a.Class)
	marshaled = binary.BigEndian.AppendUint32(marshaled, a.TTL)
	marshaled = binary.BigEndian.AppendUint16(marshaled, a.Length)
	marshaled = append(marshaled, a.Data...)

	return marshaled
}
