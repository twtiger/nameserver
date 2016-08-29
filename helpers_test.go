package nameserver

import "strings"

func createBytesForHeaders() []byte {
	return []byte{4, 210, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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

const idNum uint16 = 1234