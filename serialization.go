package nameserver

import (
	"encoding/binary"
	"errors"
)

const headerLength = 12

func (m *message) deserialize(b []byte) error {
	rem, err := extractHeaders(b)
	if err != nil {
		return err
	}
	labels, _, err := extractLabels(rem)
	if err != nil {
		return err
	}
	m.query = &query{qname: labels}
	return nil
}

func (m *message) serialize() ([]byte, error) {
	h := serializeHeaders()
	q, err := serializeQuery(m.query)
	if err != nil {
		return nil, err
	}
	b := append(h, q...)

	return b, nil
}

func serializeQuery(q *query) ([]byte, error) {
	l, err := serializeLabels(q.qname)
	if err != nil {
		return nil, err
	}
	qt := serializeUint16(uint16(q.qtype))
	qc := serializeUint16(uint16(q.qclass))

	b := append(l, append(qt, qc...)...)
	return b, nil
}

func serializeLabels(l []label) ([]byte, error) {
	if len(l) == 0 {
		return nil, errors.New("no labels to serialize")
	}

	var b []byte
	for _, e := range l {
		b = append(b, byte(len(e)))
		b = append(b, []byte(e)...)
	}
	b = append(b, 0)
	return b, nil
}

func serializeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func serializeAnswer(r []*record) []byte {
	var b []byte
	b = append(b, []byte(r[0].Name)...)

	j := make([]byte, 2)
	binary.BigEndian.PutUint16(j, uint16(r[0].Type))
	b = append(b, j...)

	k := make([]byte, 2)
	binary.BigEndian.PutUint16(k, uint16(r[0].Class))
	b = append(b, k...)

	l := make([]byte, 4)
	binary.BigEndian.PutUint32(l, uint32(r[0].TTL))
	b = append(b, l...)

	b = append(b, []byte(r[0].RData)...)
	return b
}

func extractLabels(b []byte) (l []label, remaining []byte, err error) {

	if b[0] == 0 {
		return nil, nil, errors.New("no question to extract")
	}

	for b[0] != 0 {
		length := b[0]
		lab := label(string(b[1 : length+1]))
		l = append(l, lab)
		b = b[length+1:]
	}

	return l, b[1:], nil // don't return the remaining null byte
}

func extractHeaders(in []byte) ([]byte, error) {
	if len(in) < headerLength {
		return nil, errors.New("missing header fields")
	}
	return in[headerLength:], nil
}

func serializeHeaders() []byte {
	return make([]byte, 12)
}
