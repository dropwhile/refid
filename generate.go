package refid

import (
	"crypto/rand"
	"io"
	"time"
)

func generate() ([]byte, error) {
	/*
		 0                   1                   2                   3
		 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                           unix_ts_µs                          |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                   unix_ts_µs                  |      tag      |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_b                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		|                             rand_b                            |
		+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
		7 bytes unix_ts_µs:
			 48 bits of microseconds from 1970 (about 2280 or so years worth)
		1 byte tag:
		 	255 separate tags (0 being untagged)
		8 bytes random pad:
			fill with crypto/rand random
	*/

	b := make([]byte, size)
	setTime(b, time.Now().UTC().UnixMicro())
	// use cyrpto/rand for non-test code
	err := setRandom(b, rand.Reader)
	return b, err
}

func setTime(b []byte, micros int64) {
	ms := uint64(micros)
	z := b[timeStart:]
	// A 56 bit timestamp of microseconds since epoch.
	// Which should result in about 2283 years worth of timestamps
	// 1-7 bytes: big-endian unsigned number of Unix epoch timestamp
	z[0] = byte(ms >> 48)
	z[1] = byte(ms >> 40)
	z[2] = byte(ms >> 32)
	z[3] = byte(ms >> 24)
	z[4] = byte(ms >> 16)
	z[5] = byte(ms >> 8)
	z[6] = byte(ms)
}

func setRandom(b []byte, randR io.Reader) error {
	_, err := io.ReadFull(randR, b[randStart:])
	if err != nil {
		return err
	}
	return nil
}
