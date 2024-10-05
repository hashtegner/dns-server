package dns

type Reply struct {
	Header   *Header
	Question *Question
	Answer   *Answer
}

func (m *Reply) Marshal() []byte {
	marshaled := make([]byte, 0)
	marshaled = append(marshaled, m.Header.Marshal()...)
	marshaled = append(marshaled, m.Question.Marshal()...)
	marshaled = append(marshaled, m.Answer.Marshal()...)

	return marshaled
}
