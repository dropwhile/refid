// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

var (
	refTagTest     = byte(1)
	testValWoutTag = "065f5e3p0gk013cvyvm171gn9m"
	testValWithTag = "065f5ef3q4k03p495rqw0g92sr"
)

func TestParseVarious(t *testing.T) {
	//// time prefix types
	// no tag
	_, err := Parse("065f5e3p0gk013cvyvm171gn9m")
	assert.NilError(t, err)
	_, err = Parse("018af2b8760426008d9bf6e81386154d")
	assert.NilError(t, err)
	_, err = Parse("AYryuHYEJgCNm_boE4YVTQ")
	assert.NilError(t, err)
	// with tag
	_, err = Parse("065f5ef3q4k03p495rqw0g92sr")
	assert.NilError(t, err)
	_, err = Parse("018af2b9e3b92601d8892e2fc04122ce")
	assert.NilError(t, err)
	_, err = Parse("AYryueO5JgHYiS4vwEEizg")
	assert.NilError(t, err)
	//// random time prefix types
	// no tag
	_, err = Parse("sqpwgp85q3sg1jftqhyefasemc")
	assert.NilError(t, err)
	_, err = Parse("cdedc85905b8f300c9fabc7ce7ab2ea3")
	assert.NilError(t, err)
	_, err = Parse("ze3IWQW48wDJ-rx856suow")
	assert.NilError(t, err)
	// with tag
	_, err = Parse("e02ddgb2zkyg3rmeza7jvwnn6g")
	assert.NilError(t, err)
	_, err = Parse("7004d6c162fcfd01e28efa8f2df2b534")
	assert.NilError(t, err)
	_, err = Parse("cATWwWL8_QHijvqPLfK1NA")
	assert.NilError(t, err)

	// bad ones
	_, err = Parse("nope")
	assert.Assert(t, err != nil, "expected to fail parsing invalid refid")
	_, err = Parse("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	assert.Assert(t, err != nil, "expected to fail parsing invalid refid")
	_, err = Parse("!!!!!!!!!!!!!!!!!!!!!!")
	assert.Assert(t, err != nil, "expected to fail parsing invalid refid")
	_, err = Parse("!!!!!!!!!!!!!!!!!!!!!!!!!!")
	assert.Assert(t, err != nil, "expected to fail parsing invalid refid")

	maxTime := time.UnixMilli(281474976710655)
	minTime := time.UnixMilli(0)

	// max value with type set to TimePrefix
	x, err := Parse("zzzzzzzzzzzfzzzzzzzzzzzzzw")
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0xff)
	assert.Assert(t, x.Time().Equal(maxTime))

	// max value with type set to RandomPrefix
	x, err = Parse("zzzzzzzzzzzzzzzzzzzzzzzzzw")
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0xff)
	assert.Assert(t, x.Time().Equal(minTime))

	// base32 padding at the end, so >w is truncated to w
	// TimePrefix
	x, err = Parse("zzzzzzzzzzzfzzzzzzzzzzzzzz")
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0xff)
	assert.Assert(t, x.Time().Equal(maxTime))

	// base32 padding at the end, so >w is truncated to w
	// RandomPrefix
	x, err = Parse("zzzzzzzzzzzzzzzzzzzzzzzzzz")
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0xff)
	// RandomPrefix has zero time value
	assert.Assert(t, x.Time().Equal(minTime))

	// min value with type set to TimePrefix
	x, err = Parse("00000000000000000000000000")
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0x00)
	assert.Assert(t, x.Time().Equal(minTime))

	val_max := [16]byte{}
	val_min := [16]byte{}
	for i := 0; i < 16; i++ {
		val_max[i] = 0xff
		val_min[i] = 0x00
	}

	x, err = FromBytes(val_max[:])
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0xff)
	assert.Assert(t, x.Type() == RandomPrefix)
	assert.Assert(t, x.Time().Equal(minTime))

	x, err = FromBytes(val_min[:])
	assert.NilError(t, err)
	assert.Assert(t, x.Tag() == 0x00)
	assert.Assert(t, x.Type() == TimePrefix)
	assert.Assert(t, x.Time().Equal(minTime))
}

func TestGetTime(t *testing.T) {
	t.Parallel()

	// divide times by 10, so we are close enough
	t0 := time.Now().UTC().Unix() / 10
	r := Must(New())
	vz := r.Time().UTC().Unix() / 10
	assert.Equal(t, t0, vz)

	r2 := Must(Parse(testValWoutTag))
	ts, _ := time.Parse(time.RFC3339, "2023-10-02T23:28:09.732Z")
	assert.Equal(t, ts.UTC(), r2.Time().UTC())
}

func TestSetTime(t *testing.T) {
	t.Parallel()

	ts, _ := time.Parse(time.RFC3339, "2023-01-14T18:29:00Z")

	r := Must(New())
	r.SetTime(ts)
	assert.Equal(t, ts.UTC(), r.Time().UTC())

	r = Must(NewRandom())
	// should error with random type
	err := r.SetTime(ts)
	assert.Assert(t, err != nil, "expected to error")
	assert.Assert(t, r.Time().Equal(time.UnixMilli(0)))
}

func TestBase64RoundTrip(t *testing.T) {
	t.Parallel()

	r := Must(Parse(testValWithTag))
	b64 := r.ToBase64String()
	r2, err := Parse(b64)
	assert.NilError(t, err)
	assert.Equal(t, r.String(), r2.String())
}

func TestHexRoundTrip(t *testing.T) {
	t.Parallel()

	r := Must(Parse(testValWithTag))
	b64 := r.ToHexString()
	r2, err := Parse(b64)
	assert.NilError(t, err)
	assert.Equal(t, r.String(), r2.String())
}

func TestRoundTrip(t *testing.T) {
	t.Parallel()
	u := Must(New())
	r := Must(Parse(u.String()))
	assert.Check(t, !u.HasTag(refTagTest))
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = Must(NewTagged(refTagTest))
	r = Must(Parse(u.String()))
	assert.Check(t, u.HasTag(refTagTest))
	assert.Check(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = Must(NewRandom())
	r = Must(Parse(u.String()))
	assert.Check(t, !u.HasTag(refTagTest))
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = Must(NewRandomTagged(refTagTest))
	r = Must(Parse(u.String()))
	assert.Check(t, u.HasTag(refTagTest))
	assert.Check(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())
}

func TestSetTag(t *testing.T) {
	t.Parallel()

	r := Must(Parse(testValWoutTag))
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWoutTag)
	assert.Equal(t, (&r).String(), testValWoutTag)

	r.SetTag(refTagTest)
	assert.Check(t, r.HasTag(refTagTest))
	assert.Equal(t, r.String(), "065f5e3p0gk033cvyvm171gn9m")
	assert.Equal(t, (&r).String(), "065f5e3p0gk033cvyvm171gn9m")

	r.ClearTag()
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWoutTag)
	assert.Equal(t, (&r).String(), testValWoutTag)

	r2 := Must(Parse(testValWoutTag))
	r2.SetTag(1)
	assert.Equal(t, r2.ToHexString(), "018af2b8760426018d9bf6e81386154d")
	r2.ClearTag()
	assert.Equal(t, r2.ToHexString(), "018af2b8760426008d9bf6e81386154d")
	r2.SetTag(2)
	assert.Equal(t, r2.ToHexString(), "018af2b8760426028d9bf6e81386154d")
}

func TestAmbiguous(t *testing.T) {
	t.Parallel()

	rd0 := Must(Parse(testValWoutTag))
	rd1 := Must(Parse(testValWoutTag))
	rd2 := Must(Parse(testValWoutTag))
	assert.Assert(t,
		rd0.String() == rd1.String() && rd1.String() == rd2.String(),
	)
}

func TestTemplateStringer(t *testing.T) {
	t.Parallel()
	s := Must(Parse(testValWoutTag))
	assert.Equal(t, fmt.Sprintf("%s", s), testValWoutTag)
	tpl := template.Must(template.New("name").Parse(`{{.}}`))
	var b bytes.Buffer
	err := tpl.Execute(&b, s)
	assert.NilError(t, err)
	assert.Equal(t, b.String(), testValWoutTag)
}

func TestJsonMarshal(t *testing.T) {
	t.Parallel()

	s := Must(Parse(testValWoutTag))
	j, err := json.Marshal(s)
	assert.NilError(t, err)
	assert.Equal(t, string(j), fmt.Sprintf("%q", s.String()))
}

func TestJsonUnmarshal(t *testing.T) {
	t.Parallel()

	data := fmt.Sprintf("%q", testValWoutTag)
	var r RefID
	err := json.Unmarshal([]byte(data), &r)
	assert.NilError(t, err)
	assert.Equal(t, r.String(), testValWoutTag)
}
