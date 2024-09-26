package dns

import "encoding/binary"

type Question struct {
	Labels []string
	Type   uint16 // 2 bytes (A, NS, CNAME, etc.)
	Class  uint16 // 2 bytes (IN, CH, etc.)
}

func (q *Question) Marshal() []byte {
	marshaled := make([]byte, 0)
	marshaled = append(marshaled, q.marshalLabels()...)
	marshaled = binary.BigEndian.AppendUint16(marshaled, q.Type)
	marshaled = binary.BigEndian.AppendUint16(marshaled, q.Class)

	return marshaled
}

func (q *Question) marshalLabels() []byte {
	marshaled := make([]byte, 0)
	for _, label := range q.Labels {
		marshaled = append(marshaled, byte(len(label)))
		marshaled = append(marshaled, []byte(label)...)
	}

	marshaled = append(marshaled, []byte{0}...)

	return marshaled
}
