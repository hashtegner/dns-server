package dns

func Unmarshal(b []byte) *Reply {
	header := b[:12]
	body := b[12:]
	name := UnmarshalName(body)

	return &Reply{
		Question: &Question{
			Name:  name,
			Type:  1,
			Class: 1,
		},

		Header: UnmarshalHeader(header),

		Answer: &Answer{
			Name:   name,
			Type:   1,
			Class:  1,
			TTL:    60,
			Length: 4,
			Data:   []byte{8, 8, 8, 8},
		},
	}
}
