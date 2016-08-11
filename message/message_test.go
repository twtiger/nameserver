package message

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MessageSuite struct{}

var _ = Suite(&MessageSuite{})

func (s *MessageSuite) TestMessageContainsAResourceRecordForAQuery(c *C) {
	query := Query("www.thoughtworks.com")

	response := Response(query)

	c.Assert(response.Answers, DeepEquals, []*Record{&Record{Name: "thoughtworks.com", Type: uint16(1), Class: 1, TTL: 300, RDLength: 0, RData: "161.47.4.2"}})
}
