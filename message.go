package nameserver

type label string

type header struct {
	id uint16
}

type query struct {
	qname  []label
	qtype  qType
	qclass qClass
}

type record struct {
	Name     []label
	Type     qType
	Class    qClass
	TTL      uint32
	RDLength uint16
	RData    []byte
}

type message struct {
	header  *header
	query   *query
	answers []*record
}

func (m *message) response() *message {
	return &message{
		header:  m.header,
		query:   m.query,
		answers: retrieve(m.query.qname),
	}
}
