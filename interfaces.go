package nameserver

type dnsMessage interface {
	deserialize(b []byte) error
	serialize() ([]byte, error)
}
