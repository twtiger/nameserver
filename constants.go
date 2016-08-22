package nameserver

// every type is a valid qtype
type qType uint16

// every class is a valid qclass
type qClass uint16

const (
	qtypeA qType = 1
)

// List of DNS qclass constants
const (
	qclassIN qClass = 1
)
