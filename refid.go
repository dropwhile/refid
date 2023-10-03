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
	alphabet      = "0123456789abcdefghjkmnpqrstvwxyz"
	base32Encoder = base32.NewEncoding(alphabet).WithPadding(base32.NoPadding)
	// Nil is the nil RefID, that has all 128 bits set to zero.
	Nil = RefID{}
)

const (
	size      = 16 // expected size in bytes
	typeIndex = 6  // offset where type byte starts (only lowest bit)
	tagIndex  = 7  // offset where tag starts
)

// A RefID is a 16 byte identifier that has:
//   - tagging support (support for 255 distinct tag types)
//   - go/sql scanner/valuer support
//   - multiple encodings supported: native (base32), base64, base16 (hex)
type RefID [size]byte

// New returns a new [TimePrefixed] type [RefID].
//
// If random bytes cannot be generated, it will return an error.
func New() (RefID, error) {
	var r RefID
	b, err := generate(TimePrefixed)
	if err != nil {
		return r, err
	}
	copy(r[:], b[:])
	return r, nil
}

// NewRandom returns a new [RandomPrefixed] type [RefID].
//
// If random bytes cannot be generated, it will return an error.
func NewRandom() (RefID, error) {
	var r RefID
	b, err := generate(RandomPrefixed)
	if err != nil {
		return r, err
	}
	copy(r[:], b[:])
	return r, nil
}

// NewTagged returns a new [TimePrefixed] type [RefID] tagged with tag.
//
// If random bytes cannot be generated, it will return an error.
func NewTagged(tag byte) (RefID, error) {
	r, err := New()
	if err != nil {
		return r, err
	}
	r.SetTag(tag)
	return r, nil
}

// NewRandomTagged returns a new [RandomPrefixed] type [RefID] tagged with tag.
//
// If random bytes cannot be generated, it will return an error.
func NewRandomTagged(tag byte) (RefID, error) {
	r, err := NewRandom()
	if err != nil {
		return r, err
	}
	r.SetTag(tag)
	return r, nil
}

// Parse parses a textual RefID representation, and returns
// a [RefID]. Supports parsing the following text formats:
//
//   - native/base32 (Crockford's alphabet)
//   - base64
//   - base16/hex
//
// Will return an error on parse failure.
func Parse(s string) (RefID, error) {
	var r RefID
	err := r.UnmarshalText([]byte(s))
	return r, err
}

// ParseWithRequire parses a textual RefID representation
// (same formats as Parse), while additionally requiring
// each reqs [Requirement] to pass, and returns
// a [RefID].
//
// Returns an error if RefID fails to parse or if any of the
// reqs Requirements fail.
//
// Example:
//
//	ParseWithRequire("afd661f4f2tg2vr3dca92qp6k8", HasType(RandomPrefix))
func ParseWithRequire(s string, reqs ...Requirement) (RefID, error) {
	r, err := Parse(s)
	if err != nil {
		return r, err
	}

	for _, f := range reqs {
		err = f(r)
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

// FromBytes creates a new [RefID] from a byte slice.
// Returns an error if the slice does not have a length of 16.
// The bytes are copied from the slice.
func FromBytes(input []byte) (RefID, error) {
	var r RefID
	err := r.UnmarshalBinary(input)
	return r, err
}

// FromString is an alias of [Parse].
func FromString(s string) (RefID, error) {
	return Parse(s)
}

// SetTime sets the time component of a RefID to the time
// specified by ts.
func (r *RefID) SetTime(ts time.Time) error {
	// if Radom type, do not set time, just return
	if r.HasType(RandomPrefixed) {
		return fmt.Errorf("cant set time of RandomPrefix type")
	}
	setTime(r[:], ts.UTC().UnixMilli())
	return nil
}

// Time returns the timestamp portion of a [RefID] as a [time.Time]
func (r RefID) Time() time.Time {
	if r.HasType(RandomPrefixed) {
		// if Random prefix, we have no time, so just
		// return the zero time
		return time.UnixMilli(0)
	}
	u := r[:]
	t := 0 |
		(int64(u[0]) << 40) |
		(int64(u[1]) << 32) |
		(int64(u[2]) << 24) |
		(int64(u[3]) << 16) |
		(int64(u[4]) << 8) |
		int64(u[5])
	return time.UnixMilli(t).UTC()
}

// SetTag sets the [RefID] tag to the specified value.
func (r *RefID) SetTag(tag byte) *RefID {
	r[tagIndex] = tag
	return r
}

// ClearTag clears the [RefID] tag.
func (r *RefID) ClearTag() *RefID {
	r[tagIndex] = 0
	return r
}

// IsTagged reports whether the [RefID] is tagged.
func (r RefID) IsTagged() bool {
	return r[tagIndex] != 0
}

// IsTagged reports whether the [RefID] is tagged and
// if so, if it is tagged with tag.
func (r RefID) HasTag(tag byte) bool {
	return (r.IsTagged() && r[tagIndex] == tag)
}

// Tag returns the current tag of the RefID.
// If the [RefID] is untagged, it will retrun 0.
func (r RefID) Tag() byte {
	return r[tagIndex]
}

// HasType reports whether the [RefId] is of type t.
func (r RefID) HasType(t Type) bool {
	return r[typeIndex]&0x01 == byte(t)
}

// Type returns the type of the RefID.
func (r RefID) Type() Type {
	return Type(r[typeIndex] & 0x01)
}

// IsNil reports if the [RefID] is the nil value RefID.
func (r RefID) IsNil() bool {
	return r == Nil
}

// Equal compares a [RefID] to another RefID to see
// if they have the same underlying bytes.
func (r RefID) Equal(other RefID) bool {
	return r.String() == other.String()
}

// Bytes returns a slice of a copy of the current [RefID] underlying data.
func (r RefID) Bytes() []byte {
	b := make([]byte, size)
	copy(b[:], r[:])
	return b
}

// String returns the native (base32 w/Crockford alphabet) textual represenation
// of a [RefID]
func (r RefID) String() string {
	return base32Encoder.EncodeToString(r[:])
}

// ToString is an alias of [String]
func (r RefID) ToString() string {
	return r.String()
}

// ToBase32String is an alias of [String]
func (r RefID) ToBase32String() string {
	return r.String()
}

// String returns the base64 textual represenation of a [RefID]
func (r RefID) ToBase64String() string {
	return base64.RawURLEncoding.EncodeToString(r[:])
}

// String returns the base16/hex textual represenation of a [RefID]
func (r RefID) ToHexString() string {
	return hex.EncodeToString(r[:])
}

// Format implements the [fmt.Formatter] interface.
func (r RefID) Format(f fmt.State, c rune) {
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
		fmt.Fprintf(f, "%%!%c(refid.RefID=%s)", c, r.String())
	}
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (r RefID) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// It will return an error if the slice isn't of appropriate size.
func (r *RefID) UnmarshalText(b []byte) error {
	bx := make([]byte, size)
	switch len(b) {
	case 26: // native
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
		n, err := base32Encoder.Decode(bx, b)
		if err != nil {
			return err
		}
		if n != size {
			return fmt.Errorf("wrong unmarshal size")
		}
	case 22: // base64
		n, err := base64.RawURLEncoding.Decode(bx, b)
		if err != nil {
			return err
		}
		if n != size {
			return fmt.Errorf("wrong unmarshal size")
		}
	case 32: // hex
		n, err := hex.Decode(bx, b)
		if err != nil {
			return err
		}
		if n != size {
			return fmt.Errorf("wrong unmarshal size")
		}
	default:
		return fmt.Errorf("parse error: incorrect size")
	}

	err := r.UnmarshalBinary(bx)
	if err != nil {
		return err
	}
	return nil
}

// MarshalBinary implements the [encoding.BinaryMarshaler] interface.
func (r RefID) MarshalBinary() ([]byte, error) {
	return r.Bytes(), nil
}

// UnmarshalBinary implements the [encoding.BinaryUnmarshaler] interface.
// It will return an error if the slice isn't of appropriate size.
func (r *RefID) UnmarshalBinary(data []byte) error {
	dlen := len(data)
	if dlen != size {
		return fmt.Errorf("refid: RefID must be exactly %d bytes long, got %d bytes", size, dlen)
	}
	copy(r[:], data[:])
	return nil
}

// MarshalJson implements the [json.Marshaler] interface.
func (r RefID) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// UnmarshalJson implements the [json.Unmarshaler] interface.
func (r *RefID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}
