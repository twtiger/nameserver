package nameserver

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
