package nameserver

import (
	. "gopkg.in/check.v1"
)

type DeserializationSuite struct{}

var _ = Suite(&DeserializationSuite{})

func (s *DeserializationSuite) Test_extractHeaders_returnsSliceWithoutHeaders(c *C) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2}
	rest, _ := extractHeaders(b)
	c.Assert(rest, DeepEquals, []byte{1, 2})
}

func (s *DeserializationSuite) Test_extractHeaders_returnsErrorWhenGivenSliceIsTooSmall(c *C) {
	_, e := extractHeaders([]byte{1, 2, 3})
	c.Assert(e, ErrorMatches, "missing header fields")
}

func (s *DeserializationSuite) Test_extractLabels_canParseSingleLabel(c *C) {
	b := []byte{3, byte('w'), byte('w'), byte('w'), 0, 0, 1, 0, 13}

	labels, remaining, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(labels[0], Equals, label("www"))
	c.Assert(labels[1], Equals, label(""))
	c.Assert(len(remaining), Equals, 4)
}

func (s *DeserializationSuite) Test_extractLabels_returnsRemainingBytes(c *C) {
	b := []byte{3, byte('w'), byte('w'), byte('w'), 0, 0, 1, 0, 13}

	_, remaining, err := extractLabels(b)

	c.Assert(err, IsNil)
	c.Assert(len(remaining), Equals, 4)
}

func (s *DeserializationSuite) Test_extractLabels_canParseMoreThanOneLabel(c *C) {
	b := createBytesForLabels(createLabelsFor("www.thoughtworks.com."))

	labels, _, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(len(labels), Equals, 4)
	c.Assert(labels[0], Equals, label("www"))
	c.Assert(labels[1], Equals, label("thoughtworks"))
	c.Assert(labels[2], Equals, label("com"))
	c.Assert(labels[3], Equals, label(""))
}

func (s *DeserializationSuite) Test_extractLabels_forEmptyQuestionReturnsError(c *C) {
	b := []byte{0}
	_, _, err := extractLabels(b)

	c.Assert(err, ErrorMatches, "no question to extract")
}

func (s *DeserializationSuite) Test_deserialize_returnsFullMessage(c *C) {
	b := flattenBytes(createBytesForHeaders(), 3, "www", 0, oneInTwoBytes, oneInTwoBytes)

	msg := &message{}
	err := msg.deserialize(b)

	c.Assert(err, IsNil)
	c.Assert(msg.header.id, Equals, idNum)
	c.Assert(msg.query.qname[0], Equals, label("www"))
	c.Assert(msg.query.qtype, Equals, qtypeA)
	c.Assert(msg.query.qclass, Equals, qclassIN)
}

func (s *DeserializationSuite) Test_deserialize_onQuery_returnsQuery(c *C) {
	b := flattenBytes(3, "www", 0, oneInTwoBytes, oneInTwoBytes)

	q := &query{}
	err := q.deserialize(b)

	c.Assert(err, IsNil)
	c.Assert(q.qname[0], Equals, label("www"))
	c.Assert(q.qtype, Equals, qtypeA)
	c.Assert(q.qclass, Equals, qclassIN)
}

func (s *DeserializationSuite) Test_deserialize_returnsTwoLabelsForSingleQueryAndStopsParsingAfterNullLabel(c *C) {
	b := flattenBytes(createBytesForHeaders(), 3, "www", 0, oneInTwoBytes, oneInTwoBytes, 1, 2, 3)

	msg := &message{}
	err := msg.deserialize(b)

	c.Assert(err, IsNil)
	c.Assert(len(msg.query.qname), Equals, 2)
}

func (s *DeserializationSuite) Test_deserialize_returnsErrorIfHeadersAreInvalid(c *C) {
	b := make([]byte, 1)

	msg := &message{}
	err := msg.deserialize(b)
	c.Assert(err, Not(IsNil))
}

func (s *DeserializationSuite) Test_deserialize_returnsErrorQueryIsInvalid(c *C) {
	b := make([]byte, 13)

	msg := &message{}
	err := msg.deserialize(b)
	c.Assert(err, Not(IsNil))
}

func (s *DeserializationSuite) Test_deserialize_headersCorrectly(c *C) {
	largeNumOfQueries := uint16(258)
	b := flattenBytes(idNum, 3, 4, largeNumOfQueries, 5, 6, 7, 8, 9, 10)

	h := &header{}
	h.deserialize(b)

	c.Assert(h.id, Equals, idNum)
	c.Assert(h.qdCount, Equals, largeNumOfQueries)
}
