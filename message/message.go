package message

import "strings"

// Message represents DNS messages
type Message struct {
	header  *header
	queries []*query
	answers []*Record
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
	label string
}

const a uint16 = 1
const in uint16 = 1

// Query creates a DNS query based on a given domain string
func Query(d string) *Message {
	header := &header{
		ID:      1234,
		QR:      0,
		OPCODE:  4,
		AA:      0,
		TR:      0,
		RD:      0,
		RA:      0,
		Z:       0,
		AD:      0,
		CD:      0,
		RCODE:   make([]byte, 4),
		QDCOUNT: 1,
		ANCOUNT: 0,
		NSCOUNT: 0,
		ARCOUNT: 0,
	}
	return &Message{
		header: header,
		queries: []*query{
			&query{name: &qname{labels: domainNameToLabels(d)}, qtype: a, class: in},
		},
	}
}

// Response returns the message with resource records
func Response(query *Message) *Message {
	query.answers = append(query.answers, &Record{
		Name:     "thoughtworks.com.",
		Type:     a,
		Class:    1,
		TTL:      300,
		RDLength: 0,
		RData:    "161.47.4.2",
	})
	return query
}

func domainNameToLabels(domain string) []label {
	var ls []label
	for _, v := range strings.Split(domain, ".") {
		l := &label{
			len:   uint8(len(v)),
			label: v,
		}
		ls = append(ls, *l)
	}
	return ls
}
