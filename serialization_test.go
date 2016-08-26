package nameserver

import (
	. "gopkg.in/check.v1"
)

type SerializationSuite struct{}

var _ = Suite(&SerializationSuite{})

func createBytesForAnswer() []byte {
	return flattenBytes(twTigerInBytes, oneInTwoBytes(), oneInTwoBytes(), 0, 0, 14, 16, uint16(4), 123, 123, 7, 8)
}

func createBytesForMultipleAnswers() []byte {
	answer2 := flattenBytes(twTigerInBytes, oneInTwoBytes(), oneInTwoBytes(), 0, 0, 14, 16, uint16(4), 78, 78, 90, 1)
	return append(createBytesForAnswer(), answer2...)
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForSingleLabel(c *C) {
	labels := createLabelsFor("www")

	exp := createBytesForLabels(labels)

	b, err := serializeLabels(labels)

	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_onLabel_returnsByteArray(c *C) {
	l := label("www")

	exp := flattenBytes(3, "www")
	b := l.serialize()

	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForMultipleLabels(c *C) {
	exp := twTigerInBytes

	b, err := serializeLabels(twTigerInLabels)

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

	exp := flattenBytes(twTigerInBytes, oneInTwoBytes(), oneInTwoBytes())

	q := &query{
		qname:  twTigerInLabels,
		qtype:  qtypeA,
		qclass: qclassIN,
	}

	b, _ := q.serialize()
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_forRecord_returnsByteArrayForSingleRecord(c *C) {
	record := tigerRecord1

	exp := createBytesForAnswer()

	b := record.serialize()
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeAnswer_returnsByteArrayForMultipleRecords(c *C) {
	records := []*record{
		tigerRecord1,
		tigerRecord2,
	}

	exp := createBytesForMultipleAnswers()

	b := serializeAnswer(records)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeAnswer_returnsEmptyByteArrayForNoAnswers(c *C) {
	records := []*record{}

	var exp []byte

	b := serializeAnswer(records)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeHeaders_returnsByteArrayofLength12(c *C) {
	b := serializeHeaders()
	c.Assert(len(b), Equals, 12)
}

func (s *SerializationSuite) Test_serialize_returnsByteArrayForMessageWithQuery(c *C) {
	exp := flattenBytes(createBytesForHeaders(), twTigerInBytes, oneInTwoBytes(), oneInTwoBytes())

	m := &message{
		query: &query{
			qname:  twTigerInLabels,
			qtype:  qtypeA,
			qclass: qclassIN,
		},
	}

	b, err := m.serialize()
	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_returnsByteArrayForMessageWithResponse(c *C) {
	exp := flattenBytes(createBytesForHeaders(), twTigerInBytes, oneInTwoBytes(), oneInTwoBytes(), createBytesForAnswer())

	m := &message{
		query: &query{
			qname:  twTigerInLabels,
			qtype:  qtypeA,
			qclass: qclassIN,
		},
		answers: []*record{
			tigerRecord1,
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
