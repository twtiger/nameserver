package nameserver

import . "gopkg.in/check.v1"

type SerializationSuite struct{}

var _ = Suite(&SerializationSuite{})

func createBytesForLabels() []byte {
	i := [][]byte{[]byte{3}, []byte("www"), []byte{12}, []byte("thoughtworks"), []byte{3}, []byte("com"), []byte{0}}
	return flattenBytes(i)
}

func createBytesForAnswer() []byte {
	inputs := [][]byte{[]byte("twtiger.com."), []byte{0, 1}, []byte{0, 1}, []byte{0, 0, 0, 1}, []byte("123.123.7.8")}
	return flattenBytes(inputs)
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
	i := [][]byte{createBytesForLabels(), oneInTwoBytes(), oneInTwoBytes()}

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
	i := [][]byte{createBytesForHeaders(), createBytesForLabels(), oneInTwoBytes(), oneInTwoBytes()}
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
	i := [][]byte{createBytesForHeaders(), createBytesForLabels(), oneInTwoBytes(), oneInTwoBytes(), createBytesForAnswer()}
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
