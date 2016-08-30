package nameserver

import (
	"net"
	"strings"
)

const idNum uint16 = 1234
const nsPort = 8899

var twTigerInLabels = createLabelsFor("twtiger.com")
var twTigerInBytes = createBytesForLabels(twTigerInLabels)
var ns Nameserver
var oneInTwoBytes = []byte{0, 1}

func createBytesForHeaders() []byte {
	return []byte{4, 210, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
}

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

func localhost(port int) *net.UDPAddr {
	return &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}
}

func localServer(valid bool) Nameserver {
	if !valid {
		return Nameserver{Addr: localhost(-1)}
	}
	return Nameserver{Addr: localhost(nsPort)}
}
