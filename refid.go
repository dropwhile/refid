// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

import (
	"bytes"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

var (
	// crockford base32
	// ref: https://en.wikipedia.org/wiki/Base32#Crockford's_Base32
	Alphabet         = "0123456789abcdefghjkmnpqrstvwxyz"
	WordSafeEncoding = base32.NewEncoding(Alphabet).WithPadding(base32.NoPadding)
	// Nil is the nil RefId, that has all 128 bits set to zero.
	Nil = RefId{}
)

const (
	size      = 16 // expected size in bytes
	timeStart = 0  // offset where timestamp starts
	tagIndex  = 7  // offset where tag starts
	randStart = 8  // offset where rand_b starts
)

type RefId [size]byte

// New returns a new RefId.
//
// If random bytes cannot be generated, it will return an error.
func New() (RefId, error) {
	var r RefId
	b, err := generate()
	if err != nil {
		return r, err
	}
	copy(r[:], b[:])
	return r, nil
}

// NewTagged returns a RefId tagged with tag.
//
// If random bytes cannot be generated, it will return an error.
func NewTagged(tag byte) (RefId, error) {
	r, err := New()
	if err != nil {
		return r, err
	}
	r.SetTag(tag)
	return r, nil
}

// Parse parses a textual RefId representation, and returns
// a RefId. Supports parsing the following text formats:
// * native - base32 (Crockford's alphabet)
// * base64
// * base16/hex
//
// Will return an error on parse failure.
func Parse(s string) (RefId, error) {
	var r RefId
	var err error
	switch len(s) {
	case 26: // native
		err = r.UnmarshalText([]byte(s))
	case 22: // base64
		r, err = FromBase64String(s)
	case 32: // hex
		r, err = FromHexString(s)
	default:
		return r, fmt.Errorf("parse error: incorrect size")
	}
	return r, err
}

// ParseTagged parses a textual RefId representation
// (same formats as Parse) while additionally requiring
// the parsed RefId to be tagged with tag.
//
// Returns an error if RefId fails to parse or if RefId
// is not tagged with tag.
func ParseTagged(tag byte, s string) (RefId, error) {
	r, err := Parse(s)
	if err != nil {
		return r, err
	}

	if !r.HasTag(tag) {
		return r, fmt.Errorf("RefId tag mismatch: %d != %d", r[tagIndex], tag)
	}
	return r, nil
}

// FromBytes creates a new RefId from a byte slice.
// Returns an error if the slice does not have a length of 16.
// The bytes are copied from the slice.
func FromBytes(input []byte) (RefId, error) {
	var r RefId
	err := r.UnmarshalBinary(input)
	if err != nil {
		return r, err
	}
	return r, nil
}

// FromString is an alias of Parse.
func FromString(s string) (RefId, error) {
	return Parse(s)
}

// FromBase64String parses a base64 string and returns
// a RefId.
// Returns an error if the base64 string is of improper size
// or otherwise fails to parse.
func FromBase64String(input string) (RefId, error) {
	var r RefId
	bx, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		return r, err
	}
	if len(bx) != size {
		return r, fmt.Errorf("wrong unmarshal size")
	}
	copy(r[:], bx[:])
	return r, nil
}

// FromHexString parses a base16/hex string and returns
// a RefId.
// Returns an error if the base16/hex string is of improper size
// or otherwise fails to parse.
func FromHexString(input string) (RefId, error) {
	var r RefId
	bx, err := hex.DecodeString(input)
	if err != nil {
		return r, err
	}
	if len(bx) != size {
		return r, fmt.Errorf("wrong unmarshal size")
	}
	copy(r[:], bx[:])
	return r, nil
}

// SetTime sets the time component of a RefId to the time
// specified by ts.
func (r *RefId) SetTime(ts time.Time) *RefId {
	setTime(r[:], ts.UTC().UnixMicro())
	return r
}

// Time returns the timestamp portion of a RefId as a time.Time
func (r RefId) Time() time.Time {
	u := r[timeStart:]
	t := 0 |
		(int64(u[0]) << 48) |
		(int64(u[1]) << 40) |
		(int64(u[2]) << 32) |
		(int64(u[3]) << 24) |
		(int64(u[4]) << 16) |
		(int64(u[5]) << 8) |
		int64(u[6])
	return time.UnixMicro(t).UTC()
}

// SetTag sets the RefId tag to the specified value.
func (r *RefId) SetTag(tag byte) *RefId {
	r[tagIndex] = tag
	return r
}

// ClearTag clears the RefId tag.
func (r *RefId) ClearTag() *RefId {
	r[tagIndex] = 0
	return r
}

// IsTagged reports whether the RefId is tagged.
func (r RefId) IsTagged() bool {
	return r[tagIndex] != 0
}

// IsTagged reports whether the RefId is tagged and
// if so, if it is tagged with tag.
func (r RefId) HasTag(tag byte) bool {
	return (r.IsTagged() && r[tagIndex] == tag)
}

// Tag returns the current tag of the RefId.
// If the RefId is untagged, it will retrun 0.
func (r RefId) Tag() byte {
	return r[tagIndex]
}

// IsNil reports if the RefId is the nil value RefId.
func (r RefId) IsNil() bool {
	return r == Nil
}

// Equal compares a RefId to another RefId to see
// if they have the same underlying bytes.
func (r RefId) Equal(other RefId) bool {
	return r.String() == other.String()
}

// Bytes returns a slice of a copy of the current RefId underlying data.
func (r RefId) Bytes() []byte {
	b := make([]byte, size)
	copy(b[:], r[:])
	return b
}

// String returns the native (base32 w/Crockford alphabet) textual represenation
// of a RefId
func (r RefId) String() string {
	return WordSafeEncoding.EncodeToString(r[:])
}

// ToString is an alias of String
func (r RefId) ToString() string {
	return r.String()
}

// String returns the base64 textual represenation of a RefId
func (r RefId) ToBase64String() string {
	return base64.RawURLEncoding.EncodeToString(r[:])
}

// String returns the base16/hex textual represenation of a RefId
func (r RefId) ToHexString() string {
	return hex.EncodeToString(r[:])
}

// Format implements the fmt.Formatter interface.
func (r RefId) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('#') {
		fmt.Fprintf(f, "%#v", r.Bytes())
		return
	}
	switch c {
	case 'x', 'X':
		b := make([]byte, size*2)
		hex.Encode(b, r.Bytes())
		if c == 'X' {
			bytes.ToUpper(b)
		}
		_, _ = f.Write(b)
	case 'v', 's', 'S':
		b, _ := r.MarshalText()
		if c == 'S' {
			bytes.ToUpper(b)
		}
		_, _ = f.Write(b)
	case 'q':
		_, _ = f.Write([]byte{'"'})
		_, _ = f.Write(r.Bytes())
		_, _ = f.Write([]byte{'"'})
	default:
		// invalid/unsupported format verb
		fmt.Fprintf(f, "%%!%c(refid.RefId=%s)", c, r.String())
	}
}

// MarshalText implements the encoding.TextMarshaler interface.
func (r RefId) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// It will return an error if the slice isn't of appropriate size.
func (r *RefId) UnmarshalText(b []byte) error {
	decLen := WordSafeEncoding.DecodedLen(len(b))
	if decLen != size {
		return fmt.Errorf("refid: RefId must be exactly %d bytes long, got %d bytes", size, decLen)
	}

	// lowercase, then replace ambigious chars
	b = bytes.ToLower(b)
	for i := range b {
		switch b[i] {
		case 'i', 'l':
			b[i] = '1'
		case 'o', 'O':
			b[i] = '0'
		}
	}
	bx := make([]byte, size)
	n, err := WordSafeEncoding.Decode(bx, b)
	if err != nil {
		return err
	}
	if n != size {
		return fmt.Errorf("wrong unmarshal size")
	}
	copy(r[:], bx[:])
	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (r RefId) MarshalBinary() ([]byte, error) {
	return r.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It will return an error if the slice isn't of appropriate size.
func (r *RefId) UnmarshalBinary(data []byte) error {
	dlen := len(data)
	if dlen != size {
		return fmt.Errorf("refid: RefId must be exactly %d bytes long, got %d bytes", size, dlen)
	}
	copy(r[:], data[:])
	return nil
}

// MarshalJson implements the json.Marshaler interface.
func (r RefId) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJson implements the json.Unmarshaler interface.
func (r *RefId) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}
