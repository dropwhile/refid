// Copyright (C) 2013-2018 by Maxim Bublis <b@codemonkey.ru>
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
	codecTestRefId = Must(Parse("0r32b0yermw00sbjedjxe4yaz0"))
	codecTestData  = codecTestRefId.Bytes()
)

func TestSQL(t *testing.T) {
	t.Run("Value", testSQLValue)
	t.Run("Scan", func(t *testing.T) {
		t.Run("Binary", testSQLScanBinary)
		t.Run("String", testSQLScanString)
		t.Run("Text", testSQLScanText)
		t.Run("Unsupported", testSQLScanUnsupported)
		t.Run("Nil", testSQLScanNil)
	})
}

func testSQLValue(t *testing.T) {
	v, err := codecTestRefId.Value()
	if err != nil {
		t.Fatal(err)
	}
	got, ok := v.([]byte)
	if !ok {
		t.Fatalf("Value() returned %T, want []byte", v)
	}
	want := codecTestRefId.Bytes()
	assert.Assert(
		t, cmp.DeepEqual(got, want),
		"Vlaue() == %q, want %q",
		got, want,
	)
}

func testSQLScanBinary(t *testing.T) {
	got := RefId{}
	err := got.Scan(codecTestData)
	if err != nil {
		t.Fatal(err)
	}
	if got != codecTestRefId {
		t.Errorf("Scan(%x): got %v, want %v", codecTestData, got, codecTestRefId)
	}
}

func testSQLScanString(t *testing.T) {
	s := "0r32b0yermw00sbjedjxe4yaz0"
	got := RefId{}
	err := got.Scan(s)
	if err != nil {
		t.Fatal(err)
	}
	if got != codecTestRefId {
		t.Errorf("Scan(%q): got %v, want %v", s, got, codecTestRefId)
	}
}

func testSQLScanText(t *testing.T) {
	text := []byte("0r32b0yermw00sbjedjxe4yaz0")
	got := RefId{}
	err := got.Scan(text)
	if err != nil {
		t.Fatal(err)
	}
	if got != codecTestRefId {
		t.Errorf("Scan(%q): got %v, want %v", text, got, codecTestRefId)
	}
}

func testSQLScanUnsupported(t *testing.T) {
	unsupported := []interface{}{
		true,
		42,
	}
	for _, v := range unsupported {
		got := RefId{}
		err := got.Scan(v)
		if err == nil {
			t.Errorf("Scan(%T) succeeded, got %v", v, got)
		}
	}
}

func testSQLScanNil(t *testing.T) {
	got := RefId{}
	err := got.Scan(nil)
	if err == nil {
		t.Errorf("Scan(nil) succeeded, got %v", got)
	}
}

func TestNullRefId(t *testing.T) {
	t.Run("Value", func(t *testing.T) {
		t.Run("Nil", testNullRefIdValueNil)
		t.Run("Valid", testNullRefIdValueValid)
	})

	t.Run("Scan", func(t *testing.T) {
		t.Run("Nil", testNullRefIdScanNil)
		t.Run("Valid", testNullRefIdScanValid)
		t.Run("RefId", testNullRefIdScanRefId)
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		t.Run("Nil", testNullRefIdMarshalJSONNil)
		t.Run("Null", testNullRefIdMarshalJSONNull)
		t.Run("Valid", testNullRefIdMarshalJSONValid)
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Run("Nil", testNullRefIdUnmarshalJSONNil)
		t.Run("Null", testNullRefIdUnmarshalJSONNull)
		t.Run("Valid", testNullRefIdUnmarshalJSONValid)
		t.Run("Malformed", testNullRefIdUnmarshalJSONMalformed)
	})
}

func testNullRefIdValueNil(t *testing.T) {
	nu := NullRefId{}
	got, err := nu.Value()
	if got != nil {
		t.Errorf("null NullRefId.Value returned non-nil driver.Value")
	}
	if err != nil {
		t.Errorf("null NullRefId.Value returned non-nil error")
	}
}

func testNullRefIdValueValid(t *testing.T) {
	nu := NullRefId{
		Valid: true,
		RefId: codecTestRefId,
	}
	got, err := nu.Value()
	if err != nil {
		t.Fatal(err)
	}
	s, ok := got.([]byte)
	if !ok {
		t.Errorf("Value() returned %T, want []byte", got)
	}
	want := codecTestRefId.Bytes()
	assert.Assert(
		t, cmp.DeepEqual(s, want),
		"%v.Value() == %s, want %s", nu, s, want,
	)
}

func testNullRefIdScanNil(t *testing.T) {
	u := NullRefId{}
	err := u.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}
	if u.Valid {
		t.Error("NullRefId is valid after Scan(nil)")
	}
	if u.RefId != Nil {
		t.Errorf("NullRefId.RefId is %v after Scan(nil) want Nil", u.RefId)
	}
}

func testNullRefIdScanValid(t *testing.T) {
	s := "0r32b0yermw00sbjedjxe4yaz0"
	u := NullRefId{}
	err := u.Scan(s)
	if err != nil {
		t.Fatal(err)
	}
	if !u.Valid {
		t.Errorf("Valid == false after Scan(%q)", s)
	}
	if u.RefId != codecTestRefId {
		t.Errorf("RefId == %v after Scan(%q), want %v", u.RefId, s, codecTestRefId)
	}
}

func testNullRefIdScanRefId(t *testing.T) {
	u := NullRefId{}
	err := u.Scan(codecTestRefId)
	if err != nil {
		t.Fatal(err)
	}
	if !u.Valid {
		t.Errorf("Valid == false after scan(%v)", codecTestRefId)
	}
	if u.RefId != codecTestRefId {
		t.Errorf("RefId == %v after Scan(%v), want %v", u.RefId, codecTestRefId, codecTestRefId)
	}
}

func testNullRefIdMarshalJSONNil(t *testing.T) {
	u := NullRefId{Valid: true}

	data, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("(%#v).MarshalJSON err want: <nil>, got: %v", u, err)
	}

	assert.Assert(
		t, cmp.DeepEqual(`"`+Nil.String()+`"`, string(data)),
		"(%#v).MarshalJSON value want: %s, got: %s", u, Nil.Bytes(), data,
	)
}

func testNullRefIdMarshalJSONValid(t *testing.T) {
	u := NullRefId{
		Valid: true,
		RefId: codecTestRefId,
	}

	data, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("(%#v).MarshalJSON err want: <nil>, got: %v", u, err)
	}

	assert.Assert(
		t, cmp.DeepEqual(`"`+codecTestRefId.String()+`"`, string(data)),
		"(%#v).MarshalJSON value want: %s, got: %s", u, codecTestRefId.Bytes(), data,
	)
}

func testNullRefIdMarshalJSONNull(t *testing.T) {
	u := NullRefId{}

	data, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("(%#v).MarshalJSON err want: <nil>, got: %v", u, err)
	}

	dataStr := string(data)

	if dataStr != "null" {
		t.Fatalf("(%#v).MarshalJSON value want: %s, got: %s", u, "null", dataStr)
	}
}

func testNullRefIdUnmarshalJSONNil(t *testing.T) {
	var u NullRefId

	data := []byte(`"00000000000000000000000000"`)

	if err := json.Unmarshal(data, &u); err != nil {
		t.Fatalf("json.Unmarshal err = %v, want <nil>", err)
	}

	if !u.Valid {
		t.Fatalf("u.Valid = false, want true")
	}

	if u.RefId != Nil {
		t.Fatalf("u.RefId = %v, want %v", u.RefId, Nil)
	}
}

func testNullRefIdUnmarshalJSONNull(t *testing.T) {
	var u NullRefId

	data := []byte(`null`)

	if err := json.Unmarshal(data, &u); err != nil {
		t.Fatalf("json.Unmarshal err = %v, want <nil>", err)
	}

	if u.Valid {
		t.Fatalf("u.Valid = true, want false")
	}

	if u.RefId != Nil {
		t.Fatalf("u.RefId = %v, want %v", u.RefId, Nil)
	}
}

func testNullRefIdUnmarshalJSONValid(t *testing.T) {
	var u NullRefId

	data := []byte(`"0r32b0yermw00sbjedjxe4yaz0"`)

	if err := json.Unmarshal(data, &u); err != nil {
		t.Fatalf("json.Unmarshal err = %v, want <nil>", err)
	}

	if !u.Valid {
		t.Fatalf("u.Valid = false, want true")
	}

	if u.RefId != codecTestRefId {
		t.Fatalf("u.RefId = %v, want %v", u.RefId, Nil)
	}
}

func testNullRefIdUnmarshalJSONMalformed(t *testing.T) {
	var u NullRefId

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
		n := NullRefId{RefId: u, Valid: true}
		for i := 0; i < b.N; i++ {
			n.MarshalJSON()
		}
	})
	b.Run("Invalid", func(b *testing.B) {
		n := NullRefId{Valid: false}
		for i := 0; i < b.N; i++ {
			n.MarshalJSON()
		}
	})
}

func BenchmarkNullUnmarshalJSON(b *testing.B) {
	baseRefId, err := FromString("0r32b0yermw00sbjedjxe4yaz0")
	if err != nil {
		b.Fatal(err)
	}
	data, err := json.Marshal(&baseRefId)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("Valid", func(b *testing.B) {
		var u NullRefId
		for i := 0; i < b.N; i++ {
			u.UnmarshalJSON(data)
		}
	})
	b.Run("Invalid", func(b *testing.B) {
		invalid := []byte("null")
		var n NullRefId
		for i := 0; i < b.N; i++ {
			n.UnmarshalJSON(invalid)
		}
	})
}
