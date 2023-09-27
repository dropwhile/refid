// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

import (
	"math/rand"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestGenerateTime(t *testing.T) {
	t.Parallel()

	ts, _ := time.Parse(time.RFC3339, "2023-09-14T10:27:21.826305Z")
	micros := ts.UTC().UnixMicro()

	b := make([]byte, size)
	setTime(b, micros)
	assert.DeepEqual(t, b, []byte{
		0x05, 0x4f, 0x1f, 0x0d,
		0xf0, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	})
}

func TestGenerateRandom(t *testing.T) {
	t.Parallel()

	b := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	// use a seeded psuedorandom source for testing
	r := rand.New(rand.NewSource(99))

	err := setRandom(b[8:], r)
	assert.NilError(t, err)
	// ensure first 8 bytes remain unmodified
	assert.DeepEqual(t, b[:8], []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	})
	assert.DeepEqual(t, b[8:], []byte{
		0x75, 0xed, 0x18, 0x42,
		0x49, 0xe9, 0xbc, 0x19,
	})
}

func TestGenerateBoth(t *testing.T) {
	t.Parallel()

	b := make([]byte, size)
	// use a seeded psuedorandom source for testing
	r := rand.New(rand.NewSource(99))
	ts, _ := time.Parse(time.RFC3339, "2023-09-14T10:27:21.826305Z")
	micros := ts.UTC().UnixMicro()

	setTime(b, micros)
	err := setRandom(b, r)
	assert.NilError(t, err)

	rID, err := FromBytes(b)
	assert.NilError(t, err)

	rID.SetTag(3)

	assert.DeepEqual(t, rID.Bytes(), []byte{
		0x75, 0xed, 0x18, 0x42,
		0x49, 0xe9, 0xbc, 0x03,
		0x67, 0x5e, 0x4d, 0x1f,
		0x76, 0x62, 0x13, 0xda,
	})
}
