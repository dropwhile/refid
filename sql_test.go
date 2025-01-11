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
// ref: https://github.com/gofrs/uuid/blob/22c52c268bc0dcc0569793f5b1433db423f5a9c6/sql_test.go

package refid

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

var (
	codecTestID   = Must(Parse("0r32b0yermw00sbjedjxe4yaz0"))
	codecTestData = codecTestID.Bytes()
)

func TestSQL(t *testing.T) {
	t.Run("Value", testSQLValue)
	t.Run("Scan", func(t *testing.T) {
		t.Run("BinarySlice", testSQLScanBinarySlice)
		t.Run("BinaryArray", testSQLScanBinaryArray)
		t.Run("String", testSQLScanString)
		t.Run("Text", testSQLScanText)
		t.Run("Unsupported", testSQLScanUnsupported)
		t.Run("Nil", testSQLScanNil)
	})
}

func testSQLValue(t *testing.T) {
	v, err := codecTestID.Value()
	if err != nil {
		t.Fatal(err)
	}
	got, ok := v.([]byte)
	if !ok {
		t.Fatalf("Value() returned %T, want []byte", v)
	}
	want := codecTestID.Bytes()
	assert.Assert(
		t, cmp.DeepEqual(got, want),
		"Vlaue() == %q, want %q",
		got, want,
	)
}

func testSQLScanBinarySlice(t *testing.T) {
	got := ID{}
	err := got.Scan(codecTestData)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(codecTestID) {
		t.Errorf("Scan(%x): got %v, want %v", codecTestData, got, codecTestID)
	}
}

func testSQLScanBinaryArray(t *testing.T) {
	got := ID{}
	err := got.Scan([16]byte(codecTestData))
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(codecTestID) {
		t.Errorf("Scan(%x): got %v, want %v", codecTestData, got, codecTestID)
	}
}

func testSQLScanString(t *testing.T) {
	s := "0r32b0yermw00sbjedjxe4yaz0"
	got := ID{}
	err := got.Scan(s)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(codecTestID) {
		t.Errorf("Scan(%q): got %v, want %v", s, got, codecTestID)
	}
}

func testSQLScanText(t *testing.T) {
	text := []byte("0r32b0yermw00sbjedjxe4yaz0")
	got := ID{}
	err := got.Scan(text)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Equal(codecTestID) {
		t.Errorf("Scan(%q): got %v, want %v", text, got, codecTestID)
	}
}

func testSQLScanUnsupported(t *testing.T) {
	unsupported := []interface{}{
		true,
		42,
	}
	for _, v := range unsupported {
		got := ID{}
		err := got.Scan(v)
		if err == nil {
			t.Errorf("Scan(%T) succeeded, got %v", v, got)
		}
	}
}

func testSQLScanNil(t *testing.T) {
	got := ID{}
	err := got.Scan(nil)
	if err == nil {
		t.Errorf("Scan(nil) succeeded, got %v", got)
	}
}

func TestNullID(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		t.Run("Nil", testNullIDValueNil)
		t.Run("Valid", testNullIDValueValid)
	})

	t.Run("Scan", func(t *testing.T) {
		t.Run("Nil", testNullIDScanNil)
		t.Run("Valid", testNullIDScanValid)
		t.Run("ID", testNullIDScanID)
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		t.Run("Nil", testNullIDMarshalJSONNil)
		t.Run("Null", testNullIDMarshalJSONNull)
		t.Run("Valid", testNullIDMarshalJSONValid)
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Run("Nil", testNullIDUnmarshalJSONNil)
		t.Run("Null", testNullIDUnmarshalJSONNull)
		t.Run("Valid", testNullIDUnmarshalJSONValid)
		t.Run("Malformed", testNullIDUnmarshalJSONMalformed)
	})
}

func testNullIDValueNil(t *testing.T) {
	nu := NullID{}
	got, err := nu.Value()
	if got != nil {
		t.Errorf("null NullID.Value returned non-nil driver.Value")
	}
	if err != nil {
		t.Errorf("null NullID.Value returned non-nil error")
	}
}

func testNullIDValueValid(t *testing.T) {
	nu := NullID{
		Valid: true,
		ID:    codecTestID,
	}
	got, err := nu.Value()
	if err != nil {
		t.Fatal(err)
	}
	s, ok := got.([]byte)
	if !ok {
		t.Errorf("Value() returned %T, want []byte", got)
	}
	want := codecTestID.Bytes()
	assert.Assert(
		t, cmp.DeepEqual(s, want),
		"%v.Value() == %s, want %s", nu, s, want,
	)
}

func testNullIDScanNil(t *testing.T) {
	u := NullID{}
	err := u.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if u.Valid {
		t.Error("NullID is valid after Scan(nil)")
	}
	if !u.ID.Equal(Nil) {
		t.Errorf("NullID.ID is %v after Scan(nil) want Nil", u.ID)
	}
}

func testNullIDScanValid(t *testing.T) {
	s := "0r32b0yermw00sbjedjxe4yaz0"
	u := NullID{}
	err := u.Scan(s)
	if err != nil {
		t.Fatal(err)
	}
	if !u.Valid {
		t.Errorf("Valid == false after Scan(%q)", s)
	}
	if !u.ID.Equal(codecTestID) {
		t.Errorf("ID == %v after Scan(%q), want %v", u.ID, s, codecTestID)
	}
}

func testNullIDScanID(t *testing.T) {
	u := NullID{}
	err := u.Scan(codecTestID)
	if err != nil {
		t.Fatal(err)
	}
	if !u.Valid {
		t.Errorf("Valid == false after scan(%v)", codecTestID)
	}
	if !u.ID.Equal(codecTestID) {
		t.Errorf("ID == %v after Scan(%v), want %v", u.ID, codecTestID, codecTestID)
	}
}

func testNullIDMarshalJSONNil(t *testing.T) {
	u := NullID{Valid: true}

	data, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("(%#v).MarshalJSON err want: <nil>, got: %v", u, err)
	}

	assert.Assert(
		t, cmp.DeepEqual(`"`+Nil.String()+`"`, string(data)),
		"(%#v).MarshalJSON value want: %s, got: %s", u, Nil.Bytes(), data,
	)
}

func testNullIDMarshalJSONValid(t *testing.T) {
	u := NullID{
		Valid: true,
		ID:    codecTestID,
	}

	data, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("(%#v).MarshalJSON err want: <nil>, got: %v", u, err)
	}

	assert.Assert(
		t, cmp.DeepEqual(`"`+codecTestID.String()+`"`, string(data)),
		"(%#v).MarshalJSON value want: %s, got: %s", u, codecTestID.Bytes(), data,
	)
}

func testNullIDMarshalJSONNull(t *testing.T) {
	u := NullID{}

	data, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("(%#v).MarshalJSON err want: <nil>, got: %v", u, err)
	}

	dataStr := string(data)

	if dataStr != "null" {
		t.Fatalf("(%#v).MarshalJSON value want: %s, got: %s", u, "null", dataStr)
	}
}

func testNullIDUnmarshalJSONNil(t *testing.T) {
	var u NullID

	data := []byte(`"00000000000000000000000000"`)

	if err := json.Unmarshal(data, &u); err != nil {
		t.Fatalf("json.Unmarshal err = %v, want <nil>", err)
	}

	if !u.Valid {
		t.Fatalf("u.Valid = false, want true")
	}

	if !u.ID.Equal(Nil) {
		t.Fatalf("u.ID = %v, want %v", u.ID, Nil)
	}
}

func testNullIDUnmarshalJSONNull(t *testing.T) {
	var u NullID

	data := []byte(`null`)

	if err := json.Unmarshal(data, &u); err != nil {
		t.Fatalf("json.Unmarshal err = %v, want <nil>", err)
	}

	if u.Valid {
		t.Fatalf("u.Valid = true, want false")
	}

	if !u.ID.Equal(Nil) {
		t.Fatalf("u.ID = %v, want %v", u.ID, Nil)
	}
}

func testNullIDUnmarshalJSONValid(t *testing.T) {
	var u NullID

	data := []byte(`"0r32b0yermw00sbjedjxe4yaz0"`)

	if err := json.Unmarshal(data, &u); err != nil {
		t.Fatalf("json.Unmarshal err = %v, want <nil>", err)
	}

	if !u.Valid {
		t.Fatalf("u.Valid = false, want true")
	}

	if !u.ID.Equal(codecTestID) {
		t.Fatalf("u.ID = %v, want %v", u.ID, Nil)
	}
}

func testNullIDUnmarshalJSONMalformed(t *testing.T) {
	var u NullID

	data := []byte(`257`)

	if err := json.Unmarshal(data, &u); err == nil {
		t.Fatal("json.Unmarshal err = <nil>, want error")
	}
}

func BenchmarkNullMarshalJSON(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		u, err := FromString("0r32b0yermw00sbjedjxe4yaz0")
		if err != nil {
			b.Fatal(err)
		}
		n := NullID{ID: u, Valid: true}
		for i := 0; i < b.N; i++ {
			n.MarshalJSON()
		}
	})
	b.Run("Invalid", func(b *testing.B) {
		n := NullID{Valid: false}
		for i := 0; i < b.N; i++ {
			n.MarshalJSON()
		}
	})
}

func BenchmarkNullUnmarshalJSON(b *testing.B) {
	baseID, err := FromString("0r32b0yermw00sbjedjxe4yaz0")
	if err != nil {
		b.Fatal(err)
	}
	data, err := json.Marshal(&baseID)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("Valid", func(b *testing.B) {
		var u NullID
		for i := 0; i < b.N; i++ {
			u.UnmarshalJSON(data)
		}
	})
	b.Run("Invalid", func(b *testing.B) {
		invalid := []byte("null")
		var n NullID
		for i := 0; i < b.N; i++ {
			n.UnmarshalJSON(invalid)
		}
	})
}
