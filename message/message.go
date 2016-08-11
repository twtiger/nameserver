package message

import "strings"

// Message represents DNS messages
type Message struct {
	header  *header
	queries []*query
	Answers []*Record
}

type header struct {
	ID      uint16
	QR      byte
	OPCODE  byte
	AA      byte
	TR      byte
	RD      byte
	RA      byte
	Z       byte
	AD      byte
	CD      byte
	RCODE   []byte
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

// Record represents DNS a Resource Record
type Record struct {
	Name     string
	Type     uint16
	Class    uint16
	TTL      uint32
	RDLength uint16
	RData    string
}

type query struct {
	name  *qname
	qtype uint16
	class uint16
}

type qname struct {
	labels    []label
	nullLabel byte
}

type label struct {
	len   uint8
	chars []uint8
}

const a = uint16(1)
const in = uint16(1)

// Query creates a DNS query based on a given domain string
func Query(d string) *Message {
	header := &header{
		ID:      uint16(1234),
		QR:      byte(0),
		OPCODE:  byte(4),
		AA:      byte(0),
		TR:      byte(0),
		RD:      byte(0),
		RA:      byte(0),
		Z:       byte(0),
		AD:      byte(0),
		CD:      byte(0),
		RCODE:   make([]byte, 4),
		QDCOUNT: uint16(1),
		ANCOUNT: uint16(0),
		NSCOUNT: uint16(0),
		ARCOUNT: uint16(0),
	}
	return &Message{
		header: header,
		queries: []*query{
			&query{name: &qname{labels: domainToLabels(d)}, qtype: a, class: in},
		},
	}
}

// Response returns the message with resource records
func Response(query *Message) *Message {
	return &Message{
		Answers: []*Record{
			&Record{Name: "thoughtworks.com", Type: a, Class: uint16(1), TTL: uint32(300), RDLength: uint16(0), RData: "161.47.4.2"},
		},
	}
}

func domainToLabels(domain string) []label {
	var ls []label
	for _, v := range strings.Split(domain, ".") {
		l := &label{
			len:   uint8(len(v)),
			chars: []uint8(v),
		}
		ls = append(ls, *l)
	}
	return ls
}
