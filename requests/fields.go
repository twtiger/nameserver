package requests

// FieldName represents DNS message fields
type FieldName string

// Names for fields
const (
	ID      FieldName = "id"
	QR                = "query"
	OPCODE            = "opcode"
	AA                = "authans"
	TC                = "trunc"
	RD                = "rdesc"
	RA                = "ravail"
	Z                 = "reserved"
	RCODE             = "respcode"
	QDCOUNT           = "qdcount"
	ANCOUNT           = "ancount"
	NSCOUNT           = "nscount"
	ARCOUNT           = "arcount"
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
