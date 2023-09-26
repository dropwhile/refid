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

func New() (RefId, error) {
	var r RefId
	b, err := generate()
	if err != nil {
		return r, err
	}
	copy(r[:], b[:])
	return r, nil
}

func MustNew() RefId {
	r, err := New()
	if err != nil {
		panic(err)
	}
	return r
}

func NewTagged(tag byte) (RefId, error) {
	r, err := New()
	if err != nil {
		return r, err
	}
	r.SetTag(tag)
	return r, nil
}

func MustNewTagged(tag byte) RefId {
	r := MustNew()
	r.SetTag(tag)
	return r
}

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

func MustParse(s string) RefId {
	r, err := Parse(s)
	if err != nil {
		panic(`RefId: Parse(` + s + `): ` + err.Error())
	}
	return r
}

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

func MustParseTagged(tag byte, s string) RefId {
	r, err := ParseTagged(tag, s)
	if err != nil {
		panic(`RefId: ExpectParse(` + s + `): ` + "RefId tag mismatch")
	}
	return r
}

func FromBytes(input []byte) (RefId, error) {
	var r RefId
	err := r.UnmarshalBinary(input)
	if err != nil {
		return r, err
	}
	return r, nil
}

func FromString(s string) (RefId, error) {
	return Parse(s)
}

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

func (r *RefId) SetTime(ts time.Time) *RefId {
	setTime(r[:], ts.UTC().UnixMicro())
	return r
}

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

func (r *RefId) SetTag(tag byte) *RefId {
	r[tagIndex] = tag
	return r
}

func (r *RefId) ClearTag() *RefId {
	r[tagIndex] = 0
	return r
}

func (r RefId) IsTagged() bool {
	return r[tagIndex] != 0
}

func (r RefId) HasTag(tag byte) bool {
	return (r.IsTagged() && r[tagIndex] == tag)
}

func (r RefId) Tag() byte {
	return r[tagIndex]
}

func (r RefId) IsNil() bool {
	return r == Nil
}

func (r RefId) Equal(other RefId) bool {
	return r.String() == other.String()
}

func (r RefId) Bytes() []byte {
	b := make([]byte, size)
	copy(b[:], r[:])
	return b
}

func (r RefId) String() string {
	return WordSafeEncoding.EncodeToString(r[:])
}

func (r RefId) ToString() string {
	return r.String()
}

func (r RefId) ToBase64String() string {
	return base64.RawURLEncoding.EncodeToString(r[:])
}

func (r RefId) ToHexString() string {
	return hex.EncodeToString(r[:])
}

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

func (r RefId) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

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

func (r RefId) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *RefId) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return r.UnmarshalText([]byte(s))
}
