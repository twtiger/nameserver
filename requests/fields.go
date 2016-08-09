package requests

// FieldName represents DNS message fields
type FieldName string

// Names for fields
const (
	ID       FieldName = "id"
	QUERY              = "query"
	OPCODE             = "opcode"
	AUTHANS            = "authans"
	TRUNC              = "trunc"
	RDESC              = "rdesc"
	RAVAIL             = "ravail"
	RESPCODE           = "respcode"
	QDCOUNT            = "qdcount"
	ANCOUNT            = "ancount"
	NSCOUNT            = "nscount"
	ARCOUNT            = "arcount"
)

// HeaderFieldLengths maps DNS fields with their respective lengths
var HeaderFieldLengths = map[FieldName]uint{
	ID:       16,
	QUERY:    1,
	OPCODE:   4,
	AUTHANS:  1,
	TRUNC:    1,
	RDESC:    1,
	RAVAIL:   1,
	RESPCODE: 4,
	QDCOUNT:  16,
	ANCOUNT:  16,
	NSCOUNT:  16,
	ARCOUNT:  16,
}
