package nameserver

import "strings"

type message struct {
	header   *header
	question *query
	answers  []*record
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

type record struct {
	Name     string
	Type     qType
	Class    qClass
	TTL      uint32
	RDLength uint16
	RData    string
}

type query struct {
	qname  []label
	qtype  qType
	qclass qClass
}

type label struct {
	len   uint8
	label string
}

func createMessageFor(d string) *message {
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
	return &message{
		header: header,
		question: &query{
			qname:  domainNameToLabels(d),
			qtype:  qtypeA,
			qclass: qclassIN,
		},
	}
}

// TODO change headers as needed
// TODO add any error codes if needed
func (m *message) respond() error {
	records, _ := retrieve(m.question)
	m.answers = append(m.answers, records...)
	return nil
}

// Retrieve returns a collection of resource records for a query
func retrieve(q *query) ([]*record, error) {
	// TODO: use query to perform a database lookup
	return []*record{
		&record{Name: "thoughtworks.com.", Type: qtypeA, Class: qclassIN, TTL: 300, RDLength: 0, RData: "161.47.4.2"},
	}, nil
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
