package nameserver

import (
	"encoding/binary"
	"errors"
)

func (m *message) deserialize(b []byte) error {
	rem, err := extractHeaders(b)
	if err != nil {
		return err
	}
	m.header = &header{}
	m.header.deserialize(b)
	m.query = &query{}
	err = m.query.deserialize(rem)
	if err != nil {
		return err
	}
	return nil
}

func (h *header) deserialize(b []byte) {
	h.id = binary.BigEndian.Uint16(b[:2])
	h.qdCount = binary.BigEndian.Uint16(b[4:6])
}

func (q *query) deserialize(b []byte) error {
	labels, _, err := extractLabels(b)
	if err != nil {
		return err
	}
	q.qname = labels
	q.qtype = qtypeA
	q.qclass = qclassIN
	return nil
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
