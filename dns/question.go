package dns

import "encoding/binary"

type Question struct {
	Name  *Name
	Type  uint16 // 2 bytes (A, NS, CNAME, etc.)
	Class uint16 // 2 bytes (IN, CH, etc.)
}

func (q *Question) Marshal() []byte {
	marshaled := make([]byte, 0)
	marshaled = append(marshaled, q.Name.Marshal()...)
	marshaled = binary.BigEndian.AppendUint16(marshaled, q.Type)
	marshaled = binary.BigEndian.AppendUint16(marshaled, q.Class)

	return marshaled
}
