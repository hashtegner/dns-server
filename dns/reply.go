package dns

type Reply struct {
	Header    *Header
	Questions []*Question
	Answers   []*Answer
}

func (m *Reply) Marshal() []byte {
	marshaled := make([]byte, 0)
	marshaled = append(marshaled, m.Header.Marshal()...)

	for _, question := range m.Questions {
		marshaled = append(marshaled, question.Marshal()...)

	}
	for _, answer := range m.Answers {
		marshaled = append(marshaled, answer.Marshal()...)
	}

	return marshaled
}
