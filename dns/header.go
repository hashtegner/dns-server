package dns

import (
	"encoding/binary"
)

type Header struct {
	ID    uint16 // transaction ID
	Flags *Flags // flags

	QDCount uint16 // number of entries in the question section
	AnCount uint16 // number of records in the answer section
	NsCount uint16 // number of records in the authority section
	ArCount uint16 // number of records in the additional section
}

func (m *Header) Marshal() []byte {
	marshaled := make([]byte, 12)
	binary.BigEndian.PutUint16(marshaled[0:2], m.ID)

	copy(marshaled[2:4], m.Flags.Marshal())

	binary.BigEndian.PutUint16(marshaled[4:6], m.QDCount)
	binary.BigEndian.PutUint16(marshaled[6:8], m.AnCount)

	return marshaled
}

func UnmarshalHeader(b []byte) *Header {

	return &Header{
		ID:      binary.BigEndian.Uint16(b[0:2]),
		Flags:   UnmarshalFlags(b[2:4]),
		QDCount: 1,
		AnCount: 1,
	}
}
