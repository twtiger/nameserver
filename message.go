package nameserver

type label string

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
	RData    string
}

type message struct {
	query   *query
	answers []*record
}

func (m *message) response() *message {
	return &message{
		query:   m.query,
		answers: retrieve(m.query.qname),
	}
}
