package message

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MessageSuite struct{}

var _ = Suite(&MessageSuite{})

func (s *MessageSuite) TestResourceRecordTypeAForThoughtworks(c *C) {
	message := CreateMessageFor("www.thoughtworks.com")

	message.Respond()

	c.Assert(message.question, DeepEquals, &query{
		name: &qname{
			labels: []label{
				label{len: uint8(len("www")), label: "www"},
				label{len: uint8(len("thoughtworks")), label: "thoughtworks"},
				label{len: uint8(len("com")), label: "com"},
			},
			nullLabel: 0,
		},
		qtype: a,
		class: in,
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
