package refid

//go:generate stringer -type Type
type Type byte

const (
	TimePrefixed   Type = 0x00
	RandomPrefixed Type = 0x01
)
