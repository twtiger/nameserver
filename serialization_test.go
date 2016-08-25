package nameserver

import . "gopkg.in/check.v1"

type SerializationSuite struct{}

var _ = Suite(&SerializationSuite{})

func createBytesForAnswer() []byte {
	return flattenBytes([]byte{7}, []byte("twtiger"), []byte{3}, []byte("com"), []byte{0}, []byte{0, 1}, []byte{0, 1}, []byte{0, 0, 0, 1}, []byte("123.123.7.8"))
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForSingleLabel(c *C) {
	labels := []label{label("www")}

	exp := createBytesForLabels("www")

	b, err := serializeLabels(labels)

	c.Assert(err, IsNil)
	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serialize_onLabel_returnsByteArray(c *C) {
	l := label("www")

	exp := []byte{3}
	exp = append(exp, []byte("www")...)
	b := l.serialize()

	c.Assert(b, DeepEquals, exp)
}

func (s *SerializationSuite) Test_serializeLabels_returnsByteArrayForMultipleLabels(c *C) {
	labels := []label{label("www"), label("thoughtworks"), label("com")}

	exp := createBytesForLabels("www", "thoughtworks", "com")

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
	exp := flattenBytes(createBytesForLabels("www", "thoughtworks", "com"), oneInTwoBytes(), oneInTwoBytes())

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
		Name:  []label{"twtiger", "com"},
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
			Name:  []label{"twtiger", "com"},
			Type:  1,
			Class: 1,
			TTL:   1,
			RData: "123.123.7.8",
		},
		&record{
			Name:  []label{"twtiger", "com"},
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
	exp := flattenBytes(createBytesForHeaders(), createBytesForLabels("www", "thoughtworks", "com"), oneInTwoBytes(), oneInTwoBytes())

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
	exp := flattenBytes(createBytesForHeaders(), createBytesForLabels("www", "thoughtworks", "com"), oneInTwoBytes(), oneInTwoBytes(), createBytesForAnswer())

	m := &message{
		query: &query{
			qname:  []label{label("www"), label("thoughtworks"), label("com")},
			qtype:  qtypeA,
			qclass: qclassIN,
		},
		answers: []*record{
			&record{
				Name:  []label{"twtiger", "com"},
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
