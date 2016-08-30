package nameserver

type label string

type header struct {
	id uint16
	qdCount uint16
}

type query struct {
	qname  []label
	qtype  qType
	qclass qClass
}

type record struct {
	name     []label
	_type    qType
	class    qClass
	ttl      uint32
	rdLength uint16
	rData    []byte
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
