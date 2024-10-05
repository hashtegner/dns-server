package dns

type Flags struct {
	QR     bool  // indicates whether the message is a query (0) or a reply (1)
	Opcode uint8 // 4 bits
	AA     bool  // Authoritative Answer
	TC     bool  // Truncation
	RD     bool  // Recursion Desired

	RA    bool  // Recursion Available
	Z     uint8 // 3 bits reserved for future use by DNSSEC
	Rcode uint8 // 4 bits response code, 0 => no error, 4 => not implemented
}

func UnmarshalFlags(b []byte) *Flags {
	first := b[0]
	second := b[1]
	opcode := (first & 120) >> 3

	var rcode uint8
	if opcode != 0 {
		rcode = 4
	}

	return &Flags{
		QR:     true,
		Opcode: opcode,
		AA:     (first & 4) == 1,
		TC:     (first & 2) == 1,
		RD:     (first & 1) == 1,

		RA: (second & 128) == 1,
		Z:  (second & 112) >> 4,

		Rcode: rcode,
	}
}

func (f *Flags) Marshal() []byte {

	first := byte(0)

	if f.QR {
		first |= 1 << 7
	}

	first |= f.Opcode << 3

	if f.AA {
		first |= 1 << 2
	}

	if f.TC {
		first |= 1 << 1
	}

	if f.RD {
		first |= 1
	}

	second := byte(0)

	if f.RA {
		second |= 1 << 7
	}

	second |= f.Z << 4
	second |= f.Rcode

	return []byte{first, second}
}
