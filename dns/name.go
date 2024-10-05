package dns

const nullByte = byte(0)

type Name struct {
	Labels []string
}

func UnmarshalName(b []byte) *Name {
	labels := make([]string, 0)

	cursor := byte(0)
	for true {
		length := b[cursor]
		if length == 0 {
			break
		}

		start := cursor + 1
		end := cursor + length + 1
		label := b[start:end]

		labels = append(labels, string(label))
		cursor = end
	}

	return &Name{Labels: labels}
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
