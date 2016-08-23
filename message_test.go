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

func createMessageFor(d string) *message {
	return &message{
		header: &header{},
		query: &query{
			qname:  domainNameToLabels(d),
			qtype:  qtypeA,
			qclass: qclassIN,
		},
	}
}

func (s *MessageSuite) TestResourceRecordTypeAForThoughtworks(c *C) {
	message := createMessageFor("www.thoughtworks.com")

	message.respond()

	c.Assert(message.query, DeepEquals, &query{
		qname:  []label{"www", "thoughtworks", "com"},
		qtype:  qtypeA,
		qclass: qclassIN,
	})

	c.Assert(message.answers, DeepEquals, []*record{
		&record{
			Name:     "thoughtworks.com.",
			Type:     1,
			Class:    1,
			TTL:      300,
			RDLength: 0,
			RData:    "161.47.4.2",
		},
	})
}
