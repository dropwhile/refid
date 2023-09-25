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
	testValWoutTag = "0r2nbq0wqhjg186167t0gcd1gw"
	testValWithTag = "0r2nbq0wqhjg386167t0gcd1gw"
)

func TestGetTime(t *testing.T) {
	t.Parallel()

	// divide times by 10, so we are close enough
	t0 := time.Now().UTC().Unix() / 10
	r := MustNew()
	vz := r.Time().UTC().Unix() / 10
	assert.Equal(t, t0, vz)

	r2 := MustParse(testValWoutTag)
	ts, _ := time.Parse(time.RFC3339, "2023-09-14T18:29:43.493733Z")
	assert.Equal(t, ts.UTC(), r2.Time().UTC())
}

func TestSetTime(t *testing.T) {
	t.Parallel()

	ts, _ := time.Parse(time.RFC3339, "2023-01-14T18:29:00Z")

	r := MustNew()
	r.SetTime(ts)
	assert.Equal(t, ts.UTC(), r.Time().UTC())
}

func TestBase64RoundTrip(t *testing.T) {
	t.Parallel()

	r := MustParse(testValWithTag)
	b64 := r.ToBase64String()
	r2, err := FromBase64String(b64)
	assert.NilError(t, err)
	assert.Equal(t, r.String(), r2.String())
}

func TestHexRoundTrip(t *testing.T) {
	t.Parallel()

	r := MustParse(testValWithTag)
	b64 := r.ToHexString()
	r2, err := FromHexString(b64)
	assert.NilError(t, err)
	assert.Equal(t, r.String(), r2.String())
}

func TestRoundTrip(t *testing.T) {
	t.Parallel()
	u := MustNew()
	r := MustParse(u.String())
	assert.Check(t, !u.HasTag(refTagTest))
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())

	u = MustNewTagged(refTagTest)
	r = MustParse(u.String())
	assert.Check(t, u.HasTag(refTagTest))
	assert.Check(t, r.HasTag(refTagTest))
	assert.Equal(t, u.String(), r.String())
}

func TestSetTag(t *testing.T) {
	t.Parallel()

	r := MustParse(testValWoutTag)
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWoutTag)
	assert.Equal(t, (&r).String(), testValWoutTag)

	r.SetTag(refTagTest)
	assert.Check(t, r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWithTag)
	assert.Equal(t, (&r).String(), testValWithTag)

	r.ClearTag()
	assert.Check(t, !r.HasTag(refTagTest))
	assert.Equal(t, r.String(), testValWoutTag)
	assert.Equal(t, (&r).String(), testValWoutTag)

	r2 := MustParse(testValWoutTag)
	r2.SetTag(1)
	assert.Equal(t, r2.ToHexString(), "060555dc1cbc6501a0c131f40831a187")
	r2.ClearTag()
	assert.Equal(t, r2.ToHexString(), "060555dc1cbc6500a0c131f40831a187")
	r2.SetTag(2)
	assert.Equal(t, r2.ToHexString(), "060555dc1cbc6502a0c131f40831a187")
}

func TestAmbiguous(t *testing.T) {
	t.Parallel()

	rd0 := MustParse(testValWoutTag)
	rd1 := MustParse(testValWoutTag)
	rd2 := MustParse(testValWoutTag)
	assert.Assert(t,
		rd0.String() == rd1.String() && rd1.String() == rd2.String(),
	)
}

func TestTemplateStringer(t *testing.T) {
	t.Parallel()
	s := MustParse(testValWoutTag)
	assert.Equal(t, fmt.Sprintf("%s", s), testValWoutTag)
	tpl := template.Must(template.New("name").Parse(`{{.}}`))
	var b bytes.Buffer
	err := tpl.Execute(&b, s)
	assert.NilError(t, err)
	assert.Equal(t, b.String(), testValWoutTag)
}

func TestJsonMarshal(t *testing.T) {
	t.Parallel()

	s := MustParse(testValWoutTag)
	j, err := json.Marshal(s)
	assert.NilError(t, err)
	assert.Equal(t, string(j), fmt.Sprintf("%q", s.String()))
}

func TestJsonUnmarshal(t *testing.T) {
	t.Parallel()

	data := fmt.Sprintf("%q", testValWoutTag)
	var r RefId
	err := json.Unmarshal([]byte(data), &r)
	assert.NilError(t, err)
	assert.Equal(t, r.String(), testValWoutTag)
}
