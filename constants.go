package nameserver

type qType uint16

type qClass uint16

const (
	qtypeA qType = 1
)

const (
	qclassIN qClass = 1
)

const oneHour = 3600
