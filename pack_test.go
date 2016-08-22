package nameserver

import (
	. "gopkg.in/check.v1"
)

type PackSuite struct{}

var _ = Suite(&PackSuite{})

func (s *PackSuite) Test_extractHeaders_returnsSliceWithoutHeaders(c *C) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2}
	rem, _ := extractHeaders(b)
	c.Assert(rem, DeepEquals, []byte{1, 2})
}

func (s *PackSuite) Test_extractHeaders_returnsErrorWhenGivenSliceIsTooSmall(c *C) {
	_, e := extractHeaders([]byte{1, 2, 3})
	c.Assert(e, Not(IsNil))
}

func (s *PackSuite) Test_extractLabels_canParseSingleLabel(c *C) {
	b := []byte{3, byte('w'), byte('w'), byte('w'), 0, 0, 1, 0, 13}

	labels, remaining, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(labels[0].len, Equals, uint8(3))
	c.Assert(labels[0].label, Equals, "www")
	c.Assert(len(remaining), Equals, 4)
}

func (s *PackSuite) Test_extractLabels_returnsRemainingBytes(c *C) {
	b := []byte{3, byte('w'), byte('w'), byte('w'), 0, 0, 1, 0, 13}

	_, remaining, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(len(remaining), Equals, 4)
}

func (s *PackSuite) Test_extractLabels_canParseMoreThanOneLabel(c *C) {
	b := []byte{3}
	b = append(b, []byte("www")...)
	b = append(b, 12)
	b = append(b, []byte("thoughtworks")...)
	b = append(b, 3)
	b = append(b, []byte("com")...)
	b = append(b, []byte{0, 0, 1, 3, 4}...)

	labels, _, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(labels[0].len, Equals, uint8(3))
	c.Assert(labels[0].label, Equals, "www")
	c.Assert(labels[1].len, Equals, uint8(12))
	c.Assert(labels[1].label, Equals, "thoughtworks")
	c.Assert(labels[2].len, Equals, uint8(3))
	c.Assert(labels[2].label, Equals, "com")
}

func (s *PackSuite) Test_extractLabels_forEmptyQuestionReturnsError(c *C) {
	b := []byte{0}

	_, _, err := extractLabels(b)

	c.Assert(err, ErrorMatches, "no question to extract")
}

func (s *PackSuite) TestUnpackQuestion(c *C) {
	c.Skip("p")
	b := []byte{3}
	b = append(b, []byte("www")...)
	b = append(b, 0)
	b = append(b, []byte{0, 1}...) // qtype A
	b = append(b, []byte{0, 1}...) // qclass IN

	query, err := extractQuestion(b)

	c.Assert(err, IsNil)
	c.Assert(query.qname[0].len, Equals, uint8(3))
	c.Assert(query.qname[0].label, Equals, "www")
	c.Assert(query.qtype, Equals, qtypeA)
	c.Assert(query.qclass, Equals, qclassIN)
}
