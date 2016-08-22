package nameserver

import "fmt"

type packer interface {
	unpack(b []byte) (*message, error)
	pack(*message) ([]byte, error)
}

type msgPacker struct{}

func (mp *msgPacker) unpack(b []byte) (*message, error) {
	nextLength := uint8(b[12]) // first byte after headers
	if nextLength == 0 {
		return nil, fmt.Errorf("No question to extract")
	}
	var labels []label
	domain := extractLabels(nextLength, labels, b[13:])

	return &message{
		question: &query{qname: domain},
	}, nil
}

func (mp *msgPacker) pack(m *message) ([]byte, error) {
	return nil, nil
}

func extractLabels(length uint8, labels []label, b []byte) []label {
	lab := label{
		len:   uint8(length),
		label: string(b[:length]),
	}
	labels = append(labels, lab)

	nextLength := uint8(b[length])

	if nextLength == 0 {
		return labels
	}
	b = b[length+1:]
	return extractLabels(nextLength, labels, b)
}
