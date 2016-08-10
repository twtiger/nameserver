package requests

// FieldName represents DNS message fields
type FieldName int

// Names for DNS fields
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

// HeaderFields maps DNS fields with their length, position, and offset
var HeaderFields = map[FieldName]Field{
	ID:      Field{length: 16, position: 0, offset: 0},
	QR:      Field{length: 1, position: 2, offset: 0},
	OPCODE:  Field{length: 4, position: 2, offset: 1},
	AA:      Field{length: 1, position: 2, offset: 4},
	TC:      Field{length: 1, position: 2, offset: 5},
	RD:      Field{length: 1, position: 2, offset: 6},
	RA:      Field{length: 1, position: 2, offset: 7},
	Z:       Field{length: 3, position: 3, offset: 0},
	RCODE:   Field{length: 4, position: 3, offset: 3},
	QDCOUNT: Field{length: 16, position: 4, offset: 0},
	ANCOUNT: Field{length: 16, position: 6, offset: 0},
	NSCOUNT: Field{length: 16, position: 8, offset: 0},
	ARCOUNT: Field{length: 16, position: 10, offset: 0},
}

type Field struct {
	length   uint
	position uint
	offset   uint
}
