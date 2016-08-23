package nameserver

type label string

type query struct {
	qname  []label
	qtype  qType
	qclass qClass
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
	query   *query
	answers []*record
}

func (m *message) response() *message {
	return &message{
		query: &query{
			qname:  []label{"twtiger", "com"},
			qtype:  qtypeA,
			qclass: qclassIN,
		},
		answers: []*record{
			&record{
				Name:  "twtiger.com.",
				Type:  qtypeA,
				Class: qclassIN,
				TTL:   oneHour,
				RData: "123.123.7.8",
			},
			&record{
				Name:  "twtiger.com.",
				Type:  qtypeA,
				Class: qclassIN,
				TTL:   oneHour,
				RData: "78.78.90.1",
			},
		},
	}
}
