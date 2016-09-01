package nameserver

import (
	. "gopkg.in/check.v1"
)

type SerializationSuite struct{}

var _ = Suite(&SerializationSuite{})

func bytesForAnswerFromIP(w, x, y, z int) []byte {
	return flattenBytes(twTigerInBytes, oneInTwoBytes, oneInTwoBytes, 0, 0, 14, 16, uint16(4), w, x, y, z)
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

	exp := flattenBytes(twTigerInBytes, oneInTwoBytes, oneInTwoBytes)

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

	exp := bytesForAnswerFromIP(123, 123, 7, 8)

	b := record.serialize()
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeAnswer_returnsByteArrayForMultipleRecords(c *C) {
	records := []*record{
		tigerRecord1,
		tigerRecord2,
	}

	exp := append(bytesForAnswerFromIP(123, 123, 7, 8), bytesForAnswerFromIP(78, 78, 90, 1)...)

	b := serializeAnswer(records)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeAnswer_returnsEmptyByteArrayForNoAnswers(c *C) {
	records := []*record{}

	var exp []byte

	b := serializeAnswer(records)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeHeaders_returnsPopulatedByteArrayofLength12(c *C) {
	header := &header{
		id:      idNum,
		qdCount: oneQuery,
		anCount: twoAnswers,
	}

	b := serializeHeaders(header)

	c.Assert(len(b), Equals, 12)
	c.Assert(b, DeepEquals, []byte{4, 210, 0, 0, 0, 1, 0, 2, 0, 0, 0, 0})
}

func (s *SerializationSuite) Test_serialize_returnsByteArrayForMessageWithQuery(c *C) {
	exp := flattenBytes(createBytesForHeaders(), twTigerInBytes, oneInTwoBytes, oneInTwoBytes)

	m := &message{
		header: &header{
			id:      idNum,
			qdCount: oneQuery,
			anCount: twoAnswers,
		},
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
	exp := flattenBytes(createBytesForHeaders(), twTigerInBytes, oneInTwoBytes, oneInTwoBytes, bytesForAnswerFromIP(123, 123, 7, 8))

	m := &message{
		header: &header{
			id:      idNum,
			qdCount: oneQuery,
			anCount: twoAnswers,
		},
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

func (s *SerializationSuite) Test_flattenDifferentTypes(c *C) {
	exp := append(append(append(append([]byte("foo"), []byte{12}...), []byte{3}...), []byte("bar")...), []byte{0, byte(123)}...)

	bytes := flattenBytes("foo", 12, byte(3), []byte("bar"), uint16(123))

	c.Assert(bytes, DeepEquals, exp)
}

func (s *SerializationSuite) Test_flattenUint16(c *C) {
	exp := []byte{1, 1}

	bytes := flattenBytes(uint16(257))

	c.Assert(bytes, DeepEquals, exp)
}

func (s *SerializationSuite) Test_flattenUint32(c *C) {
	exp := []byte{1, 0, 0, 1}

	bytes := flattenBytes(uint32(16777217))

	c.Assert(bytes, DeepEquals, exp)
}
