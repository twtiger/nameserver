package nameserver

import . "gopkg.in/check.v1"

type SerializationSuite struct{}

var _ = Suite(&SerializationSuite{})

func createBytesForHeaders() []byte {
	return []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func createBytesForLabels() []byte {
	i := [][]byte{[]byte{3}, []byte("www"), []byte{12}, []byte("thoughtworks"), []byte{3}, []byte("com"), []byte{0}}
	return flattenBytes(i)
}

func createBytesForUint16() []byte {
	return []byte{0, 1}
}

func flattenBytes(i [][]byte) (b []byte) {
	for _, e := range i {
		b = append(b, e...)
	}
	return b
}

func createBytesForAnswer() []byte {
	inputs := [][]byte{[]byte("twtiger.com."), []byte{0, 1}, []byte{0, 1}, []byte{0, 0, 0, 1}, []byte("123.123.7.8")}
	return flattenBytes(inputs)
}

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
	i := [][]byte{[]byte{3}, []byte("www"), []byte{12}, []byte("thoughtworks"), []byte{3}, []byte("com"), []byte{0}}
	b := flattenBytes(i)

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
	b := createBytesForHeaders()
	b = append(b, 3)
	b = append(b, []byte("www")...)
	b = append(b, 0)

	msg := &message{}
	err := msg.deserialize(b)

	c.Assert(err, IsNil)
	c.Assert(msg.query.qname[0], Equals, label("www"))
}

func (s *SerializationSuite) Test_deserialize_returnsOneLabelForSingleQueryAndStopsParsingAfterNullLabel(c *C) {
	b := createBytesForHeaders()
	b = append(b, []byte{3, byte('w'), byte('w'), byte('w'), 0}...)
	b = append(b, []byte{0, 0, 1, 3, 4}...)

	msg := &message{}
	err := msg.deserialize(b)

	c.Assert(err, IsNil)
	c.Assert(len(msg.query.qname), Equals, 1)
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

	exp := []byte{3, byte('w'), byte('w'), byte('w'), 0}

	b, err := serializeLabels(labels)

	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_onLabel_returnsByteArray(c *C) {
	l := label("www")

	exp := []byte{3, byte('w'), byte('w'), byte('w')}

	b := l.serialize()

	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForMultipleLabels(c *C) {
	labels := []label{label("www"), label("thoughtworks"), label("com")}

	i := [][]byte{createBytesForLabels()}
	exp := flattenBytes(i)

	b, err := serializeLabels(labels)

	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeLabels_returnsErrorForNoLabels(c *C) {
	labels := []label{}

	_, err := serializeLabels(labels)

	c.Assert(err, ErrorMatches, "no labels to serialize")
}

func (s *SerializationSuite) Test_serializeUint16_returnsByteArray(c *C) {
	exp := []byte{0, 1}
	b := serializeUint16(uint16(qtypeA))
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeUint32_returnsByteArray(c *C) {
	exp := []byte{0, 1, 0, 0}
	b := serializeUint32(uint32(65536))
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeQuery_returnsByteArrayForMessageQuery(c *C) {
	i := [][]byte{createBytesForLabels(), createBytesForUint16(), createBytesForUint16()}

	exp := flattenBytes(i)

	q := &query{
		qname:  []label{label("www"), label("thoughtworks"), label("com")},
		qtype:  qtypeA,
		qclass: qclassIN,
	}

	b, _ := q.serialize()
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_forRecord_returnsByteArrayForSingleRecord(c *C) {
	record := &record{
		Name:  "twtiger.com.",
		Type:  1,
		Class: 1,
		TTL:   1,
		RData: "123.123.7.8",
	}

	exp := createBytesForAnswer()

	b := record.serialize()
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeAnswer_returnsByteArrayForMultipleRecords(c *C) {
	records := []*record{
		&record{
			Name:  "twtiger.com.",
			Type:  1,
			Class: 1,
			TTL:   1,
			RData: "123.123.7.8",
		},
		&record{
			Name:  "twtiger.com.",
			Type:  1,
			Class: 1,
			TTL:   1,
			RData: "123.123.7.8",
		},
	}

	exp := createBytesForAnswer()
	exp = append(exp, createBytesForAnswer()...)

	b := serializeAnswer(records)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeHeaders_returnsByteArrayofLength12(c *C) {
	b := serializeHeaders()
	c.Assert(len(b), Equals, 12)
}

func (s *SerializationSuite) Test_serialize_returnsByteArrayForMessageWithQuery(c *C) {
	i := [][]byte{createBytesForHeaders(), createBytesForLabels(), createBytesForUint16(), createBytesForUint16()}
	exp := flattenBytes(i)

	m := &message{
		query: &query{
			qname:  []label{label("www"), label("thoughtworks"), label("com")},
			qtype:  qtypeA,
			qclass: qclassIN,
		},
	}

	b, err := m.serialize()
	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_returnsByteArrayForMessageWithResponse(c *C) {
	i := [][]byte{createBytesForHeaders(), createBytesForLabels(), createBytesForUint16(), createBytesForUint16(), createBytesForAnswer()}
	exp := flattenBytes(i)

	m := &message{
		query: &query{
			qname:  []label{label("www"), label("thoughtworks"), label("com")},
			qtype:  qtypeA,
			qclass: qclassIN,
		},
		answers: []*record{
			&record{
				Name:  "twtiger.com.",
				Type:  1,
				Class: 1,
				TTL:   1,
				RData: "123.123.7.8",
			},
		},
	}

	b, err := m.serialize()
	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_returnsErrorForInvalidQueryWithNoLabels(c *C) {

	m := &message{
		query: &query{
			qname: []label{},
		},
	}

	_, err := m.serialize()
	c.Assert(err, ErrorMatches, "no labels to serialize")
}
