package nameserver

import (
	"fmt"
	"strings"
)

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

var oneInTwoBytes = []byte{0, 1}

func createBytesForLabels(l []label) (b []byte) {
	for _, e := range l {
		b = flattenBytes(b, len(string(e)), string(e))
	}
	b = append(b, 0)
	return
}

func createLabelsFor(s string) (labels []label) {
	a := strings.Split(s, ".")
	for _, l := range a {
		labels = append(labels, label(l))
	}
	return
}

var twTigerInLabels = createLabelsFor("twtiger.com")
var twTigerInBytes = createBytesForLabels(twTigerInLabels)
