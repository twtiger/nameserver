package nameserver

type label string

type header struct {
	id      uint16
	qdCount uint16
	anCount uint16
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
	records := retrieve(m.query.qname)
	numOfRecords := uint16(len(records))
	return &message{
		header: &header{
			id:      m.header.id,
			qdCount: m.header.qdCount,
			anCount: numOfRecords,
		},
		query:   m.query,
		answers: records,
	}
}
