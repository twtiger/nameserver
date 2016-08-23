package nameserver

import (
	. "gopkg.in/check.v1"
)

type SerializationSuite struct{}

var _ = Suite(&SerializationSuite{})

func (s *SerializationSuite) Test_extractHeaders_returnsSliceWithoutHeaders(c *C) {
	b := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2}
	rem, _ := extractHeaders(b)
	c.Assert(rem, DeepEquals, []byte{1, 2})
}

func (s *SerializationSuite) Test_extractHeaders_returnsErrorWhenGivenSliceIsTooSmall(c *C) {
	_, e := extractHeaders([]byte{1, 2, 3})
	c.Assert(e, ErrorMatches, "missing header fields")
}

func (s *SerializationSuite) Test_extractLabels_canParseSingleLabel(c *C) {
	b := []byte{3, byte('w'), byte('w'), byte('w'), 0, 0, 1, 0, 13}

	labels, remaining, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(labels[0], Equals, label("www"))
	c.Assert(len(remaining), Equals, 4)
}

func (s *SerializationSuite) Test_extractLabels_returnsRemainingBytes(c *C) {
	b := []byte{3, byte('w'), byte('w'), byte('w'), 0, 0, 1, 0, 13}

	_, remaining, err := extractLabels(b)

	c.Assert(err, IsNil)
	c.Assert(len(remaining), Equals, 4)
}

func (s *SerializationSuite) Test_extractLabels_canParseMoreThanOneLabel(c *C) {
	b := []byte{3}
	b = append(b, []byte("www")...)
	b = append(b, 12)
	b = append(b, []byte("thoughtworks")...)
	b = append(b, 3)
	b = append(b, []byte("com")...)
	b = append(b, []byte{0, 0, 1, 3, 4}...)

	labels, _, err := extractLabels(b)

	c.Assert(err, IsNil)

	c.Assert(labels[0], Equals, label("www"))
	c.Assert(labels[1], Equals, label("thoughtworks"))
	c.Assert(labels[2], Equals, label("com"))
}

func (s *SerializationSuite) Test_extractLabels_forEmptyQuestionReturnsError(c *C) {
	b := []byte{0}
	_, _, err := extractLabels(b)

	c.Assert(err, ErrorMatches, "no question to extract")
}

func (s *SerializationSuite) Test_deserialize_returnsMessageWithQuery(c *C) {
	b := make([]byte, 12)
	b = append(b, 3)
	b = append(b, []byte("www")...)
	b = append(b, 0)
	b = append(b, []byte{0, 0, 1, 3, 4}...)

	msg := &message{}
	err := msg.deserialize(b)

	c.Assert(err, IsNil)
	c.Assert(msg.query.qname[0], Equals, label("www"))
}

func (s *SerializationSuite) Test_deserialize_returnsErrorIfHeadersAreInvalid(c *C) {
	b := make([]byte, 1)

	msg := &message{}
	err := msg.deserialize(b)
	c.Assert(err, Not(IsNil))
}

func (s *SerializationSuite) Test_deserialize_returnsErrorQueryIsInvalid(c *C) {
	b := make([]byte, 13)

	msg := &message{}
	err := msg.deserialize(b)
	c.Assert(err, Not(IsNil))
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForSingleLabel(c *C) {
	labels := []label{label("www")}

	exp := []byte{3}
	exp = append(exp, []byte("www")...)
	exp = append(exp, 0)

	b, err := serializeLabels(labels)

	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForMultipleLabels(c *C) {
	labels := []label{label("www"), label("thoughtworks"), label("com")}

	exp := []byte{3}
	exp = append(exp, []byte("www")...)
	exp = append(exp, 12)
	exp = append(exp, []byte("thoughtworks")...)
	exp = append(exp, 3)
	exp = append(exp, []byte("com")...)
	exp = append(exp, 0)

	b, err := serializeLabels(labels)

	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}
