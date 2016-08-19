package database

// Persistance is an interface for a type that is able to lookup queries for a DNS record
// TODO make better interface name
type Persistance interface {
	Lookup(name string, qType, qClass int) ([]Record, error)
}

// Database holds DNS Zones and their corresponding Records in a tree-like structure
type database struct{}

// Lookup takes a record name, type, and class, and returns a list of corresponding records and an error
func (p *database) Lookup(name string, qType, qClass int) ([]Record, error) {
	// TODO
	return nil, nil
}

// Record is a DNS record
type Record struct{}
