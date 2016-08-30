package nameserver

import (
	"strings"

	. "gopkg.in/check.v1"
)

type MessageSuite struct{}

var _ = Suite(&MessageSuite{})

func domainNameToLabels(domain string) []label {
	labels := []label{}
	for _, p := range strings.Split(domain, ".") {
		labels = append(labels, label(p))
	}
	return labels
}

func createQueryFor(d string, id uint16) *message {
	return &message{
		header: &header{
			id: id,
			qdCount: 1,
		},
		query: &query{
			qname:  domainNameToLabels(d),
			qtype:  qtypeA,
			qclass: qclassIN,
		},
	}
}

func (s *MessageSuite) Test_ResponseForAuthoritativeZoneQuery(c *C) {
	q := createQueryFor("twtiger.com", 1234)

	r := q.response()

	c.Assert(r.header, DeepEquals, &header{
		id: 1234,
		qdCount: 1,
	})
	c.Assert(r.query, DeepEquals, &query{
		qname:  []label{"twtiger", "com"},
		qtype:  qtypeA,
		qclass: qclassIN,
	})
	c.Assert(len(r.answers), Equals, 2)

	c.Assert(r.answers[0], DeepEquals,
		&record{
			name:     []label{"twtiger", "com"},
			_type:    qtypeA,
			class:    qclassIN,
			ttl:      oneHour,
			rdLength: 4,
			rData:    []byte{123, 123, 7, 8},
		})
	c.Assert(r.answers[1], DeepEquals,
		&record{
			name:     []label{"twtiger", "com"},
			_type:    qtypeA,
			class:    qclassIN,
			ttl:      oneHour,
			rdLength: 4,
			rData:    []byte{78, 78, 90, 1},
		})
}

func (s *MessageSuite) Test_ResponseForExtNameServerQuery(c *C) {
	q := createQueryFor("wireshark.org", 456)

	r := q.response()

	c.Assert(r.header, DeepEquals, &header{id: 456, qdCount: 1})

	c.Assert(r.query, DeepEquals, &query{
		qname:  []label{"wireshark", "org"},
		qtype:  qtypeA,
		qclass: qclassIN,
	})
	c.Assert(len(r.answers), Equals, 0)
}
