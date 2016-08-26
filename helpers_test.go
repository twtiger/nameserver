package nameserver

import "fmt"

func flattenBytes(i ...interface{}) (b []byte) {
	for _, e := range i {
		switch k := e.(type) {
		case string:
			b = append(b, []byte(k)...)
		case int:
			b = append(b, byte(k))
		case byte:
			b = append(b, k)
		case []byte:
			b = append(b, k...)
		case uint16:
			b = append(b, []byte{0, byte(k)}...)
		default:
			panic(fmt.Sprintf("cannot flatten: %#v", e))
		}
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
		b = flattenBytes(b, len(e), e)
	}
	b = append(b, 0)
	return
}
