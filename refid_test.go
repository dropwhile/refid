package refid

import (
	"bytes"
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
	refId := MustNew()
	vz := refId.Time().UTC().Unix() / 10
	assert.Equal(t, t0, vz)

	refId2 := MustParse(testValWoutTag)
	ts, _ := time.Parse(time.RFC3339, "2023-09-14T18:29:43.493733Z")
	assert.Equal(t, ts.UTC(), refId2.Time().UTC())
}

func TestSetTime(t *testing.T) {
	t.Parallel()

	ts, _ := time.Parse(time.RFC3339, "2023-01-14T18:29:00Z")

	refId := MustNew()
	refId.SetTime(ts)
	assert.Equal(t, ts.UTC(), refId.Time().UTC())
}

func TestBase64RoundTrip(t *testing.T) {
	t.Parallel()

	refId := MustParse(testValWithTag)
	b64 := refId.ToBase64String()
	refId2, err := FromBase64String(b64)
	assert.NilError(t, err)
	assert.Equal(t, refId.String(), refId2.String())
}

func TestHexRoundTrip(t *testing.T) {
	t.Parallel()

	refId := MustParse(testValWithTag)
	b64 := refId.ToHexString()
	refId2, err := FromHexString(b64)
	assert.NilError(t, err)
	assert.Equal(t, refId.String(), refId2.String())
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

	refId := MustParse(testValWoutTag)
	assert.Check(t, !refId.HasTag(refTagTest))
	assert.Equal(t, refId.String(), testValWoutTag)
	assert.Equal(t, (&refId).String(), testValWoutTag)

	refId.SetTag(refTagTest)
	assert.Check(t, refId.HasTag(refTagTest))
	assert.Equal(t, refId.String(), testValWithTag)
	assert.Equal(t, (&refId).String(), testValWithTag)

	refId.ClearTag()
	assert.Check(t, !refId.HasTag(refTagTest))
	assert.Equal(t, refId.String(), testValWoutTag)
	assert.Equal(t, (&refId).String(), testValWoutTag)

	refId3 := MustParse(testValWoutTag)
	refId3.SetTag(1)
	assert.Equal(t, refId3.ToHexString(), "060555dc1cbc6501a0c131f40831a187")
	refId3.ClearTag()
	assert.Equal(t, refId3.ToHexString(), "060555dc1cbc6500a0c131f40831a187")
	refId3.SetTag(2)
	assert.Equal(t, refId3.ToHexString(), "060555dc1cbc6502a0c131f40831a187")
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
