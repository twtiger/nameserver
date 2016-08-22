package nameserver

import . "gopkg.in/check.v1"

type MessageSuite struct{}

var _ = Suite(&MessageSuite{})

func (s *MessageSuite) TestResourceRecordTypeAForThoughtworks(c *C) {
	message := createMessageFor("www.thoughtworks.com")

	message.respond()

	c.Assert(message.question, DeepEquals, &query{
		qname: []label{
			label{len: uint8(len("www")), label: "www"},
			label{len: uint8(len("thoughtworks")), label: "thoughtworks"},
			label{len: uint8(len("com")), label: "com"},
		},
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
