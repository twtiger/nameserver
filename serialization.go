package nameserver

import "errors"

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
	l, err := serializeLabels(m.query.qname)
	if err != nil {
		return nil, err
	}
	b := append(h, l...)

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
