package requests

// FieldName represents DNS message fields
type FieldName int

// Names for fields
const (
	ID FieldName = iota
	QR
	OPCODE
	AA
	TC
	RD
	RA
	Z
	RCODE
	QDCOUNT
	ANCOUNT
	NSCOUNT
	ARCOUNT
)

// HeaderFieldLengths maps DNS fields with their respective lengths
var HeaderFieldLengths = map[FieldName]uint{
	ID:      16,
	QR:      1,
	OPCODE:  4,
	AA:      1,
	TC:      1,
	RD:      1,
	RA:      1,
	Z:       3,
	RCODE:   4,
	QDCOUNT: 16,
	ANCOUNT: 16,
	NSCOUNT: 16,
	ARCOUNT: 16,
}
