package nameserver

var tigerRecord1 = &record{
	name:     []label{"twtiger", "com", ""},
	_type:    qtypeA,
	class:    qclassIN,
	ttl:      oneHour,
	rdLength: 4,
	rData:    []byte{123, 123, 7, 8},
}

var tigerRecord2 = &record{
	name:     []label{"twtiger", "com", ""},
	_type:    qtypeA,
	class:    qclassIN,
	ttl:      oneHour,
	rdLength: 4,
	rData:    []byte{78, 78, 90, 1},
}
