package dns

import (
	"bytes"
)

const headerSize = 12
const questionFlagsSize = 4

func Unmarshal(buff []byte) *Reply {
	rawHeader := buff[:headerSize]
	rawQuestions := buff[headerSize:]

	header := UnmarshalHeader(rawHeader)
	questions := make([]*Question, header.QDCount)

	offset := 0
	for i := 0; i < int(header.QDCount); i++ {
		questions[i] = UnmarshalQuestion(rawQuestions, offset)
		questionLength := bytes.Index(rawQuestions[offset:], []byte{0}) + questionFlagsSize

		offset += questionLength + 1
	}

	header.AnCount = uint16(len(questions))
	answers := make([]*Answer, len(questions))

	for i := 0; i < len(questions); i++ {
		question := questions[i]

		answers[i] = &Answer{
			Name:   question.Name,
			Type:   question.Type,
			Class:  question.Class,
			TTL:    60,
			Length: 4,
			Data:   []byte{8, 8, 8, 8},
		}
	}

	return &Reply{
		Header:    header,
		Questions: questions,
		Answers:   answers,
	}
}
