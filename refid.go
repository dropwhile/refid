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
	var refId RefId
	b, err := generate()
	if err != nil {
		return refId, err
	}
	copy(refId[:], b[:])
	return refId, nil
}

func MustNew() RefId {
	refId, err := New()
	if err != nil {
		panic(err)
	}
	return refId
}

func NewTagged(tag byte) (RefId, error) {
	refId, err := New()
	if err != nil {
		return refId, err
	}
	refId.SetTag(tag)
	return refId, nil
}

func MustNewTagged(tag byte) RefId {
	refId := MustNew()
	refId.SetTag(tag)
	return refId
}

func Parse(s string) (RefId, error) {
	var refId RefId
	err := refId.UnmarshalText([]byte(s))
	if err != nil {
		return refId, err
	}
	return refId, nil
}

func MustParse(s string) RefId {
	refId, err := Parse(s)
	if err != nil {
		panic(`RefId: Parse(` + s + `): ` + err.Error())
	}
	return refId
}

func ParseTagged(tag byte, s string) (RefId, error) {
	refId, err := Parse(s)
	if err != nil {
		return refId, err
	}

	if !refId.HasTag(tag) {
		return refId, fmt.Errorf("RefId tag mismatch: %d != %d", refId[tagIndex], tag)
	}
	return refId, nil
}

func MustParseTagged(tag byte, s string) RefId {
	refId, err := ParseTagged(tag, s)
	if err != nil {
		panic(`RefId: ExpectParse(` + s + `): ` + "RefId tag mismatch")
	}
	return refId
}

func FromBytes(input []byte) (RefId, error) {
	var refId RefId
	err := refId.UnmarshalBinary(input)
	if err != nil {
		return refId, err
	}
	return refId, nil
}

func FromString(s string) (RefId, error) {
	return Parse(s)
}

func FromBase64String(input string) (RefId, error) {
	var refId RefId
	bx, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		return refId, err
	}
	if len(bx) != size {
		return refId, fmt.Errorf("wrong unmarshal size")
	}
	copy(refId[:], bx[:])
	return refId, nil
}

func FromHexString(input string) (RefId, error) {
	var refId RefId
	bx, err := hex.DecodeString(input)
	if err != nil {
		return refId, err
	}
	if len(bx) != size {
		return refId, fmt.Errorf("wrong unmarshal size")
	}
	copy(refId[:], bx[:])
	return refId, nil
}

func (refId *RefId) SetTime(ts time.Time) *RefId {
	setTime(refId[:], ts.UTC().UnixMicro())
	return refId
}

func (refId *RefId) SetTag(tag byte) *RefId {
	refId[tagIndex] = tag
	return refId
}

func (refId *RefId) ClearTag() *RefId {
	refId[tagIndex] = 0
	return refId
}

func (refId RefId) IsTagged() bool {
	return refId[tagIndex] != 0
}

func (refId RefId) HasTag(tag byte) bool {
	return (refId.IsTagged() && refId[tagIndex] == tag)
}

func (refId RefId) Tag() byte {
	return refId[tagIndex]
}

func (refId RefId) IsNil() bool {
	return refId == Nil
}

func (refId RefId) Equal(other RefId) bool {
	return refId.String() == other.String()
}

func (refId RefId) MarshalText() ([]byte, error) {
	return []byte(refId.String()), nil
}

func (refId RefId) Time() time.Time {
	u := refId[timeStart:]
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

func (refId *RefId) UnmarshalText(b []byte) error {
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
	copy(refId[:], bx[:])
	return nil
}

func (refId RefId) Bytes() []byte {
	b := make([]byte, size)
	copy(b[:], refId[:])
	return b
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (refId RefId) MarshalBinary() ([]byte, error) {
	return refId.Bytes(), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It will return an error if the slice isn't of appropriate size.
func (refId *RefId) UnmarshalBinary(data []byte) error {
	dlen := len(data)
	if dlen != size {
		return fmt.Errorf("refid: RefId must be exactly %d bytes long, got %d bytes", size, dlen)
	}
	copy(refId[:], data[:])
	return nil
}

func (refId RefId) String() string {
	return WordSafeEncoding.EncodeToString(refId[:])
}

func (refId RefId) ToString() string {
	return refId.String()
}

func (refId RefId) ToBase64String() string {
	return base64.RawURLEncoding.EncodeToString(refId[:])
}

func (refId RefId) ToHexString() string {
	return hex.EncodeToString(refId[:])
}

func (refId RefId) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('#') {
		fmt.Fprintf(f, "%#v", refId.Bytes())
		return
	}
	switch c {
	case 'x', 'X':
		b := make([]byte, size*2)
		hex.Encode(b, refId.Bytes())
		if c == 'X' {
			bytes.ToUpper(b)
		}
		_, _ = f.Write(b)
	case 'v', 's', 'S':
		b, _ := refId.MarshalText()
		if c == 'S' {
			bytes.ToUpper(b)
		}
		_, _ = f.Write(b)
	case 'q':
		_, _ = f.Write([]byte{'"'})
		_, _ = f.Write(refId.Bytes())
		_, _ = f.Write([]byte{'"'})
	default:
		// invalid/unsupported format verb
		fmt.Fprintf(f, "%%!%c(refid.RefId=%s)", c, refId.String())
	}
}

func (refId RefId) MarshalJSON() ([]byte, error) {
	return json.Marshal(refId.String())
}

func (refid *RefId) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return refid.UnmarshalText([]byte(s))
}
