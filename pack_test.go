package nameserver

import . "gopkg.in/check.v1"

type PackSuite struct{}

var _ = Suite(&PackSuite{})

func (s *PackSuite) TestUnpackSingleQName(c *C) {
	p := msgPacker{}
	b := make([]byte, 12)
	b = append(b, byte(3))
	b = append(b, []byte("www")...)
	b = append(b, byte(0))

	msg, err := p.unpack(b)

	c.Assert(err, IsNil)
	qname := msg.question.qname
	c.Assert(len(qname.labels[0]), Equals, 3)
	c.Assert(qname.labels[0], Equals, label("www"))
}

func (s *PackSuite) TestUnpackMultipleQNames(c *C) {
	p := msgPacker{}
	b := make([]byte, 12)
	b = append(b, byte(3))
	b = append(b, []byte("www")...)
	b = append(b, byte(12))
	b = append(b, []byte("thoughtworks")...)
	b = append(b, byte(3))
	b = append(b, []byte("com")...)
	b = append(b, byte(0))

	msg, err := p.unpack(b)

	c.Assert(err, IsNil)

	qname := msg.question.qname
	c.Assert(len(qname.labels[0]), Equals, 3)
	c.Assert(qname.labels[0], Equals, label("www"))
	c.Assert(len(qname.labels[1]), Equals, 12)
	c.Assert(qname.labels[1], Equals, label("thoughtworks"))
	c.Assert(len(qname.labels[2]), Equals, 3)
	c.Assert(qname.labels[2], Equals, label("com"))
}

func (s *PackSuite) TestUnpackEmptyQuestionReturnsError(c *C) {
	p := msgPacker{}
	b := make([]byte, 12)
	b = append(b, byte(0))

	msg, err := p.unpack(b)

	c.Assert(err, ErrorMatches, "No question to extract")
	c.Assert(msg, IsNil)
}
