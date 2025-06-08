package dns

import (
	"encoding/binary"
)

const nullByte = byte(0)

type Name struct {
	Labels []string
}

func UnmarshalName(b []byte, offset int) *Name {
	return &Name{Labels: unmarshalLabels(b, offset)}
}

func (n *Name) Marshal() []byte {
	marshaled := make([]byte, 0)

	for _, label := range n.Labels {
		marshaled = append(marshaled, byte(len(label)))
		marshaled = append(marshaled, []byte(label)...)
	}

	marshaled = append(marshaled, nullByte)

	return marshaled
}

func unmarshalLabels(buff []byte, offset int) []string {
	labels := make([]string, 0)

	for {
		length := int(buff[offset])
		if length == 0 {
			break
		}

		// dns compression
		// https://www.rfc-editor.org/rfc/rfc1035#section-4.1.4
		// the first bit is 1, so it's a pointer,
		// we need to retrieve the pointer and use it to jump
		// to the real label and then continue unmarshalling the labels
		if length&0xc0 == 0xc0 {
			pointer := buff[offset : offset+2] // 16 bits pointer

			// given the 16 bits pointer, the first 2 bits are 1, and the rest (14 bits) is the real pointer value
			newOffset := binary.BigEndian.Uint16(pointer)&0x3FFF - headerSize

			labels = append(labels, unmarshalLabels(buff, int(newOffset))...)
			offset += 2

			continue
		}

		start := offset + 1
		end := start + length
		label := buff[start:end]
		labels = append(labels, string(label))
		offset = end
	}

	return labels
}
