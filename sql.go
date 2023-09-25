package refid

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
)

var (
	_ driver.Valuer = RefId{}
	_ sql.Scanner   = (*RefId)(nil)
)

// Value implements the driver.Valuer interface.
func (u RefId) Value() (driver.Value, error) {
	return u.Bytes(), nil
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice will be handled by UnmarshalBinary, while
// a longer byte slice or a string will be handled by UnmarshalText.
func (u *RefId) Scan(src interface{}) error {
	switch src := src.(type) {
	case RefId: // support gorm convert from RefId to NullRefId
		*u = src
		return nil

	case []byte:
		if len(src) == size {
			return u.UnmarshalBinary(src)
		}
		return u.UnmarshalText(src)

	case string:
		var parseFunc func(string) (RefId, error)
		switch len(src) {
		case 26: // native
			parseFunc = Parse
		case 32: // hex
			parseFunc = FromHexString
		case 22: // base64
			parseFunc = FromBase64String
		default:
			return fmt.Errorf("refid: cannot convert %T to RefId", src)
		}
		uu, err := parseFunc(src)
		*u = uu
		return err
	}

	return fmt.Errorf("refid: cannot convert %T to RefId", src)
}

// NullRefId can be used with the standard sql package to represent a
// RefId value that can be NULL in the database.
type NullRefId struct {
	RefId RefId
	Valid bool
}

// Value implements the driver.Valuer interface.
func (u NullRefId) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	return u.RefId.Value()
}

// Scan implements the sql.Scanner interface.
func (u *NullRefId) Scan(src interface{}) error {
	if src == nil {
		u.RefId, u.Valid = Nil, false
		return nil
	}

	u.Valid = true
	return u.RefId.Scan(src)
}

var nullJSON = []byte("null")

// MarshalJSON marshals the NullRefId as null or the nested RefId
func (u NullRefId) MarshalJSON() ([]byte, error) {
	if !u.Valid {
		return nullJSON, nil
	}
	var buf bytes.Buffer
	buf.WriteByte('"')
	buf.WriteString(u.RefId.String())
	buf.WriteByte('"')
	return buf.Bytes(), nil
}

// UnmarshalJSON unmarshals a NullRefId
func (u *NullRefId) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		u.RefId, u.Valid = Nil, false
		return nil
	}
	if n := len(b); n >= 2 && b[0] == '"' {
		b = b[1 : n-1]
	}
	err := u.RefId.UnmarshalText(b)
	u.Valid = (err == nil)
	return err
}
