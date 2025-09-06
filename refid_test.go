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

	"github.com/dropwhile/assert"
)

var (
	refTagTest     = byte(1)
	testValWoutTag = "1hh39ed2w3n01fwsmk9w4d4x44"
	testValWithTag = "1hh39ed2w3n03fwsmk9w4d4x44"
)

func TestParseVarious(t *testing.T) {
	//// time prefix types
	// no tag
	_, err := Parse("065f5e3p0gk013cvyvm171gn9m")
	assert.Nil(t, err)
	_, err = Parse("018af2b8760426008d9bf6e81386154d")
	assert.Nil(t, err)
	_, err = Parse("AYryuHYEJgCNm_boE4YVTQ")
	assert.Nil(t, err)
	// with tag
	_, err = Parse("065f5ef3q4k03p495rqw0g92sr")
	assert.Nil(t, err)
	_, err = Parse("018af2b9e3b92601d8892e2fc04122ce")
	assert.Nil(t, err)
	_, err = Parse("AYryueO5JgHYiS4vwEEizg")
	assert.Nil(t, err)
	//// random time prefix types
	// no tag
	_, err = Parse("sqpwgp85q3sg1jftqhyefasemc")
	assert.Nil(t, err)
	_, err = Parse("cdedc85905b8f300c9fabc7ce7ab2ea3")
	assert.Nil(t, err)
	_, err = Parse("ze3IWQW48wDJ-rx856suow")
	assert.Nil(t, err)
	// with tag
	_, err = Parse("e02ddgb2zkyg3rmeza7jvwnn6g")
	assert.Nil(t, err)
	_, err = Parse("7004d6c162fcfd01e28efa8f2df2b534")
	assert.Nil(t, err)
	_, err = Parse("cATWwWL8_QHijvqPLfK1NA")
	assert.Nil(t, err)

	// bad ones
	_, err = Parse("nope")
	assert.NotNil(t, err, "expected to fail parsing invalid refid")
	_, err = Parse("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	assert.NotNil(t, err, "expected to fail parsing invalid refid")
	_, err = Parse("!!!!!!!!!!!!!!!!!!!!!!")
	assert.NotNil(t, err, "expected to fail parsing invalid refid")
	_, err = Parse("!!!!!!!!!!!!!!!!!!!!!!!!!!")
	assert.NotNil(t, err, "expected to fail parsing invalid refid")

	maxTime := time.UnixMilli(35184372088831)
	minTime := time.UnixMilli(0)

	// max value with type set to TimePrefixed
	x, err := Parse("zzzzzzzzzzzfzzzzzzzzzzzzzw")
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0xff)
	assert.True(t, x.Time().Equal(maxTime))

	// max value with type set to RandomPrefixed
	x, err = Parse("zzzzzzzzzzzzzzzzzzzzzzzzzw")
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0xff)
	assert.True(t, x.Time().Equal(minTime))

	// base32 padding at the end, so >w is truncated to w
	// TimePrefixed
	x, err = Parse("zzzzzzzzzzzfzzzzzzzzzzzzzz")
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0xff)
	assert.True(t, x.Time().Equal(maxTime))

	// base32 padding at the end, so >w is truncated to w
	// RandomPrefixed
	x, err = Parse("zzzzzzzzzzzzzzzzzzzzzzzzzz")
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0xff)
	// RandomPrefixed has zero time value
	assert.True(t, x.Time().Equal(minTime))

	// min value with type set to TimePrefixed
	x, err = Parse("00000000000000000000000000")
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0x00)
	assert.True(t, x.Time().Equal(minTime))

	val_max := [16]byte{}
	val_min := [16]byte{}
	for i := 0; i < 16; i++ {
		val_max[i] = 0xff
		val_min[i] = 0x00
	}

	x, err = FromBytes(val_max[:])
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0xff)
	assert.True(t, x.Type() == RandomPrefixed)
	assert.True(t, x.Time().Equal(minTime))

	x, err = FromBytes(val_min[:])
	assert.Nil(t, err)
	assert.True(t, x.Tag() == 0x00)
	assert.True(t, x.Type() == TimePrefixed)
	assert.True(t, x.Time().Equal(minTime))
}

func TestGetTime(t *testing.T) {
	t.Parallel()

	// divide times by 10, so we are close enough
	t0 := time.Now().UTC().Unix() / 10
	r := Must(New())
	vz := r.Time().UTC().Unix() / 10
	assert.Equal(t, t0, vz)

	r2 := Must(Parse(testValWoutTag))
	fmt.Println(r2.Time())
	ts, _ := time.Parse(time.RFC3339, "2023-12-07T23:22:43.676Z")
	assert.Equal(t, ts.UTC(), r2.Time().UTC())
}

func TestSetTime(t *testing.T) {
	t.Parallel()

	ts, _ := time.Parse(time.RFC3339, "2023-01-14T18:29:00Z")

	r := Must(New())
	r.SetTime(ts)
	assert.Equal(t, ts.UTC(), r.Time().UTC())

	// try to set a time too far into the future
	err := r.SetTime(time.UnixMilli(maxTime + 1))
	fmt.Printf("ERRRR: %s\n", err)
	assert.NotNil(t, err, "expected to error")

	r = Must(NewRandom())
	// should error with random type
	err = r.SetTime(ts)
	assert.NotNil(t, err, "expected to error")
	assert.True(t, r.Time().Equal(time.UnixMilli(0)))
}

func TestBase64RoundTrip(t *testing.T) {
	t.Parallel()

	r := Must(Parse(testValWithTag))
	b64 := r.ToBase64String()
	r2, err := Parse(b64)
	assert.Nil(t, err)
	assert.Equal(t, r.String(), r2.String())
}

func TestHexRoundTrip(t *testing.T) {
	t.Parallel()

	r := Must(Parse(testValWithTag))
	b64 := r.ToHexString()
	r2, err := Parse(b64)
	assert.Nil(t, err)
	assert.Equal(t, r.String(), r2.String())
}

func TestRoundTrip(t *testing.T) {
	t.Parallel()
	u := Must(New())
	r := Must(Parse(u.String()))
	assert.False(t, u.HasTag(refTagTest))
	assert.False(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = Must(NewTagged(refTagTest))
	r = Must(Parse(u.String()))
	assert.True(t, u.HasTag(refTagTest))
	assert.True(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = Must(NewRandom())
	r = Must(Parse(u.String()))
	assert.False(t, u.HasTag(refTagTest))
	assert.False(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = Must(NewRandomTagged(refTagTest))
	r = Must(Parse(u.String()))
	assert.True(t, u.HasTag(refTagTest))
	assert.True(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())
}

func TestSetTag(t *testing.T) {
	t.Parallel()

	r := Must(Parse(testValWoutTag))
	assert.False(t, r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWoutTag)
	assert.Equal(t, (&r).String(), testValWoutTag)

	r.SetTag(refTagTest)
	assert.True(t, r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWithTag)
	assert.Equal(t, (&r).String(), testValWithTag)

	r.ClearTag()
	assert.False(t, r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWoutTag)
	assert.Equal(t, (&r).String(), testValWoutTag)

	r2 := Must(Parse(testValWoutTag))
	r2.SetTag(1)
	assert.Equal(t, r2.ToHexString(), "0c6234b9a2e0ea01bf99a4d3c2349d21")
	r2.ClearTag()
	assert.Equal(t, r2.ToHexString(), "0c6234b9a2e0ea00bf99a4d3c2349d21")
	r2.SetTag(2)
	assert.Equal(t, r2.ToHexString(), "0c6234b9a2e0ea02bf99a4d3c2349d21")
}

func TestAmbiguous(t *testing.T) {
	t.Parallel()

	rd0 := Must(Parse(testValWoutTag))
	rd1 := Must(Parse(testValWoutTag))
	rd2 := Must(Parse(testValWoutTag))
	assert.True(t, rd0.String() == rd1.String() && rd1.String() == rd2.String())
	assert.True(t, rd0.Equal(rd1) && rd1.Equal(rd2))
}

func TestTemplateStringer(t *testing.T) {
	t.Parallel()
	s := Must(Parse(testValWoutTag))
	assert.Equal(t, fmt.Sprintf("%s", s), testValWoutTag)
	tpl := template.Must(template.New("name").Parse(`{{.}}`))
	var b bytes.Buffer
	err := tpl.Execute(&b, s)
	assert.Nil(t, err)
	assert.Equal(t, b.String(), testValWoutTag)
}

func TestJsonMarshal(t *testing.T) {
	t.Parallel()

	s := Must(Parse(testValWoutTag))
	j, err := json.Marshal(s)
	assert.Nil(t, err)
	assert.Equal(t, string(j), fmt.Sprintf("%q", s.String()))
}

func TestJsonUnmarshal(t *testing.T) {
	t.Parallel()

	data := fmt.Sprintf("%q", testValWoutTag)
	var r ID
	err := json.Unmarshal([]byte(data), &r)
	assert.Nil(t, err)
	assert.Equal(t, r.String(), testValWoutTag)
}
