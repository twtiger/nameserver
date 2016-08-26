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

var database map[string][]*record

func init() {
	database = make(map[string][]*record)
	database["twtiger.com."] = []*record{tigerRecord1, tigerRecord2}
}

func retrieve(labels []label) (rrs []*record) {
	recordName := ""
	for _, l := range labels {
		recordName += string(l) + "."
	}

	if rrs, ok := database[recordName]; ok {
		return rrs
	}
	return []*record{}
}
