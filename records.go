package nameserver

var tigerRecord1 = &record{
	Name:     []label{"twtiger", "com"},
	Type:     qtypeA,
	Class:    qclassIN,
	TTL:      oneHour,
	RDLength: 4,
	RData:    []byte{123, 123, 7, 8},
}

var tigerRecord2 = &record{
	Name:     []label{"twtiger", "com"},
	Type:     qtypeA,
	Class:    qclassIN,
	TTL:      oneHour,
	RDLength: 4,
	RData:    []byte{78, 78, 90, 1},
}
