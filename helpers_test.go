package nameserver

func flattenBytes(i ...[]byte) (b []byte) {
	for _, e := range i {
		b = append(b, e...)
	}
	return b
}

func createBytesForHeaders() []byte {
	return []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func oneInTwoBytes() []byte {
	return []byte{0, 1}
}

func createBytesForLabels(s ...string) (b []byte) {
	for _, e := range s {
		b = append(b, append([]byte{byte(len(e))}, []byte(e)...)...)
	}
	b = append(b, 0)
	return
}