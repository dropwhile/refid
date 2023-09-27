// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

import (
	"crypto/rand"
	"io"
	"time"
)

func generateTimePrefixType() ([]byte, error) {
	b := make([]byte, size)
	setTime(b, time.Now().UTC().UnixMilli())
	err := setRandom(b[typeIndex:], rand.Reader)
	// clear tag
	b[tagIndex] = 0x00
	// set type to 0 (clear lowest bit)
	b[typeIndex] &^= 0x01
	return b, err
}

func generateRandomPrefixType() ([]byte, error) {
	b := make([]byte, size)
	err := setRandom(b, rand.Reader)
	// clear tag
	b[tagIndex] = 0x00
	// set type to 1 (set bit 1)
	b[typeIndex] |= 0x01
	return b, err
}

func setTime(b []byte, millis int64) {
	ms := uint64(millis)
	// A 56 bit timestamp of microseconds since epoch.
	// Which should result in about 2283 years worth of timestamps
	// 1-7 bytes: big-endian unsigned number of Unix epoch timestamp
	b[0] = byte(ms >> 40)
	b[1] = byte(ms >> 32)
	b[2] = byte(ms >> 24)
	b[3] = byte(ms >> 16)
	b[4] = byte(ms >> 8)
	b[5] = byte(ms)
}

// use cyrpto/rand for non-test code
func setRandom(b []byte, randR io.Reader) error {
	_, err := io.ReadFull(randR, b)
	if err != nil {
		return err
	}
	return nil
}
