// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

func generate(t Type) ([]byte, error) {
	b := make([]byte, size)
	var err error
	switch t {
	case TimePrefixed:
		err = setRandom(b[typeIndex-1:], rand.Reader)
		setTime(b, time.Now().UTC().UnixMilli())
		// set type to 0 (clear lowest bit)
		b[typeIndex] &^= 0x01
	case RandomPrefixed:
		err = setRandom(b, rand.Reader)
		// set type to 1 (set bit 1)
		b[typeIndex] |= 0x01
	default:
		return b, fmt.Errorf("unknown type specified")
	}
	// clear tag
	b[tagIndex] = 0x00
	return b, err
}

func setTime(b []byte, millis int64) {
	// A 45 bit timestamp of milliseconds since epoch.
	// Which should be fine until around year 3084
	// 1-7 bytes: big-endian unsigned number of Unix epoch timestamp
	ms := uint64(millis) << 3
	b[0] = byte(ms >> 40)
	b[1] = byte(ms >> 32)
	b[2] = byte(ms >> 24)
	b[3] = byte(ms >> 16)
	b[4] = byte(ms >> 8)
	// clear all but the bottom 3 bits of b[5],
	// as that is random data we want to leave as radom
	b[5] = byte(ms) | (b[5] & 0b111)
}

// use cyrpto/rand for non-test code
func setRandom(b []byte, randR io.Reader) error {
	_, err := io.ReadFull(randR, b)
	if err != nil {
		return err
	}
	return nil
}
