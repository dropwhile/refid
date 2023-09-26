// Original Work: Copyright (C) 2013-2018 by Maxim Bublis <b@codemonkey.ru>
// Modifications: Copyright (C) 2023 Eli Janssen
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// ref: https://github.com/gofrs/uuid/blob/22c52c268bc0dcc0569793f5b1433db423f5a9c6/sql.go

package refid

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	_ driver.Valuer = RefId{}
	_ sql.Scanner   = (*RefId)(nil)
)

// Value implements the [sql/driver.Valuer] interface.
func (r RefId) Value() (driver.Value, error) {
	return r.Bytes(), nil
}

// Scan implements the [sql.Scanner] interface.
// A 16-byte slice will be handled by [RefId.UnmarshalBinary], while
// a longer byte slice or a string will be handled by [RefId.UnmarshalText].
func (r *RefId) Scan(src interface{}) error {
	switch src := src.(type) {
	case RefId: // support gorm convert from RefId to NullRefId
		*r = src
		return nil

	case []byte:
		if len(src) == size {
			return r.UnmarshalBinary(src)
		}
		return r.UnmarshalText(src)

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
		*r = uu
		return err
	}

	return fmt.Errorf("refid: cannot convert %T to RefId", src)
}

// NullRefId can be used with the standard sql package to represent a
// [RefId] value that can be NULL in the database.
type NullRefId struct {
	RefId RefId
	Valid bool
}

// Value implements the [sql/driver.Valuer] interface.
func (u NullRefId) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	return u.RefId.Value()
}

// Scan implements the [sql.Scanner] interface.
func (u *NullRefId) Scan(src interface{}) error {
	if src == nil {
		u.RefId, u.Valid = Nil, false
		return nil
	}

	u.Valid = true
	return u.RefId.Scan(src)
}

var nullJSON = []byte("null")

// MarshalJSON marshals the [NullRefId] as null or the nested [RefId]
func (u NullRefId) MarshalJSON() ([]byte, error) {
	if !u.Valid {
		return nullJSON, nil
	}
	return json.Marshal(u.RefId.String())
}

// UnmarshalJSON unmarshals a [NullRefId]
func (u *NullRefId) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		u.RefId, u.Valid = Nil, false
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	err := u.RefId.UnmarshalText([]byte(s))
	u.Valid = (err == nil)
	return err
}
