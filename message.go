package nameserver

type label string

type query struct {
	qname  []label
	qtype  qType
	qclass qClass
}

// TODO change headers as needed
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

type message struct {
	header  *header
	query   *query
	answers []*record
}

// TODO add any error codes if needed
// TODO this should return a new message
func (m *message) respond() error {
	records := []*record{
		&record{Name: "thoughtworks.com.", Type: qtypeA, Class: qclassIN, TTL: 300, RDLength: 0, RData: "161.47.4.2"},
	}
	m.answers = append(m.answers, records...)
	return nil
}
