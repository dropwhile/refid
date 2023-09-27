package refid

//go:generate stringer -type Type
type Type byte

const (
	/*
		TimePrefix Type (0x00)
		 0                   1                   2                   3
		 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                           unix_ts_ms                          |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|           unix_ts_ms          |    rand_a   |t|      tag      |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_b                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_b                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		unix_ts_ms:
			48 bits big-endian unsigned number of Unix epoch timestamp milliseconds.
			(2284 years worth... until about year 4200 or so)
		rand_a:
			7 bits random pad. fill with crypto/rand random.
		t:
			1 bit for type. either RefId (type:0) or RefIdRand (type:1)
		tag:
			8 bits tag. 255 separate tags (0 being untagged).
		rand_b:
			64 bits random pad. fill with crypto/rand random.
	*/
	TimePrefix Type = 0x00
	/*
		RandomPrefix Type (0x01)
		0                   1                   2                   3
		0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_a                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                    rand_a                   |t|      tag      |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_b                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_b                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		rand_a:
			55 bits random pad. fill with crypto/rand random.
		t:
			1 bit for type. either RefId (type:0) or RefIdRand (type:1)
		tag:
			8 bits tag. 255 separate tags (0 being untagged).
		rand_b:
			64 bits random pad. fill with crypto/rand random.
	*/
	RandomPrefix Type = 0x01
)
