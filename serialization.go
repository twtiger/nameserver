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
	q, err := m.query.serialize()
	if err != nil {
		return nil, err
	}

	return append(serializeHeaders(), append(q, serializeAnswer(m.answers)...)...), nil
}

func (q *query) serialize() ([]byte, error) {
	l, err := serializeLabels(q.qname)
	if err != nil {
		return nil, err
	}
	qt := serializeUint16(uint16(q.qtype))
	qc := serializeUint16(uint16(q.qclass))

	b := append(l, append(qt, qc...)...)
	return b, nil
}

func (r *record) serialize() (b []byte) {
	b = append(b, []byte(r.Name)...)

	rt := serializeUint16(uint16(r.Type))
	b = append(b, rt...)

	rc := serializeUint16(uint16(r.Class))
	b = append(b, rc...)

	rttl := serializeUint32(uint32(r.TTL))
	b = append(b, rttl...)

	b = append(b, []byte(r.RData)...)
	return b
}

func (l label) serialize() (b []byte) {
	b = append(b, byte(len(l)))
	b = append(b, []byte(l)...)
	return b
}

func serializeLabels(l []label) (b []byte, err error) {
	if len(l) == 0 {
		return nil, errors.New("no labels to serialize")
	}

	for _, e := range l {
		b = append(b, e.serialize()...)
	}
	b = append(b, 0)
	return b, nil
}

func serializeUint16(i uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return b
}

func serializeUint32(i uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return b
}

func serializeAnswer(r []*record) (b []byte) {
	for _, e := range r {
		b = append(b, e.serialize()...)
	}
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
