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
			id:      id,
			qdCount: oneQuery,
		},
		query: &query{
			qname:  domainNameToLabels(d),
			qtype:  qtypeA,
			qclass: qclassIN,
		},
	}
}

func (s *MessageSuite) Test_ResponseForAuthoritativeZoneQuery(c *C) {
	q := createQueryFor("twtiger.com.", 1234)

	r := q.response()

	c.Assert(r.header, DeepEquals,
		&header{
			id:      1234,
			qdCount: oneQuery,
			anCount: twoAnswers,
		})
	c.Assert(r.query, DeepEquals, q.query)

	c.Assert(len(r.answers), Equals, 2)
	c.Assert(r.answers, DeepEquals,
		[]*record{
			tigerRecord1,
			tigerRecord2,
		})
}

func (s *MessageSuite) Test_ResponseForExtNameServerQuery(c *C) {
	q := createQueryFor("wireshark.org.", 456)

	r := q.response()

	c.Assert(r.header, DeepEquals,
		&header{
			id:      456,
			qdCount: oneQuery,
			anCount: 0,
		})
	c.Assert(r.query, DeepEquals, q.query)

	c.Assert(len(r.answers), Equals, 0)
}
