package dns

func Unmarshal(b []byte) *Reply {
	// _header := b[:12]
	body := b[12:]
	name := UnmarshalName(body)

	return &Reply{
		Question: &Question{
			Name:  name,
			Type:  1,
			Class: 1,
		},

		Header: &Header{
			ID:      1234,
			QDCount: 1,
			AnCount: 1,
			Flags: &Flags{
				QR: true,
			},
		},

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
