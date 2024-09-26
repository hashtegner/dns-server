package dns

type Message struct {
	Header   *Header
	Question *Question
}

func (m *Message) Marshal() []byte {
	marshaled := make([]byte, 0)
	marshaled = append(marshaled, m.Header.Marshal()...)
	marshaled = append(marshaled, m.Question.Marshal()...)

	return marshaled
}

func Unmarshal(b []byte) *Message {
	// _header := b[:12]
	body := b[12:]
	labels := parseLabels(body)

	return &Message{
		Question: &Question{
			Labels: labels,
			Type:   1,
			Class:  1,
		},

		Header: &Header{
			ID:      1234,
			QDCount: 1,
			Flags: &Flags{
				QR: true,
			},
		},
	}
}

func parseLabels(b []byte) []string {
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

	return labels
}
