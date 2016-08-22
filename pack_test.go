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
	c.Assert(qname.labels[0].len, Equals, uint8(3))
	c.Assert(qname.labels[0].label, Equals, "www")
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
	c.Assert(qname.labels[0].len, Equals, uint8(3))
	c.Assert(qname.labels[0].label, Equals, "www")
	c.Assert(qname.labels[1].len, Equals, uint8(12))
	c.Assert(qname.labels[1].label, Equals, "thoughtworks")
	c.Assert(qname.labels[2].len, Equals, uint8(3))
	c.Assert(qname.labels[2].label, Equals, "com")
}

func (s *PackSuite) TestUnpackEmptyQuestionReturnsError(c *C) {
	p := msgPacker{}
	b := make([]byte, 12)
	b = append(b, byte(0))

	msg, err := p.unpack(b)

	c.Assert(err, ErrorMatches, "No question to extract")
	c.Assert(msg, IsNil)
}
