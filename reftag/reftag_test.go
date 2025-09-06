package reftag

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/dropwhile/assert"
	"github.com/dropwhile/refid/v2"
)

func TestParse(t *testing.T) {
	r, err := Parse[IDt100]("065s6jt2bvn68rxnnaezjv5wkg")
	assert.Nil(t, err)
	assert.Equal(t,
		fmt.Sprintf("%T", r),
		"reftag.IDt100",
	)
}

func TestParseWithRequire(t *testing.T) {
	r, err := ParseWithRequire[IDt100](
		"065s6jt2bvn68rxnnaezjv5wkg",
		refid.HasType(refid.TimePrefixed),
	)
	assert.Nil(t, err)
	assert.Equal(t,
		fmt.Sprintf("%T", r),
		"reftag.IDt100",
	)
}

func TestFromBytes(t *testing.T) {
	b, err := hex.DecodeString("018b934b425eea6463b5aa9df96cbc9c")
	assert.Nil(t, err)
	r, err := FromBytes[IDt100](b)
	assert.Nil(t, err)
	assert.Equal(t,
		fmt.Sprintf("%T", r),
		"reftag.IDt100",
	)
}

func TestMatcher(t *testing.T) {
	r, err := Parse[IDt100]("065s6jt2bvn68rxnnaezjv5wkg")
	assert.Nil(t, err)
	r2, err := Parse[IDt101]("065s6xkqdfn6bscrt7a8dz2sag")
	assert.Nil(t, err)

	m := NewMatcher[IDt100]()
	assert.True(t, m.Match(r))
	assert.True(t, m.Match(&r))
	assert.True(t, m.Match(r.String()))
	assert.True(t, m.Match(r.Bytes()))
	assert.True(t, !m.Match(r2))
	assert.True(t, !m.Match(&r2))
	assert.True(t, !m.Match(r2.String()))
	assert.True(t, !m.Match(r2.Bytes()))
}

func TestTypeAliasing(t *testing.T) {
	type (
		MyID  = IDt6
		MyID2 = IDt100
	)

	var (
		NewMyID     = New[MyID]
		MyIDMatcher = NewMatcher[MyID]
		ParseMyID   = Parse[MyID]
		ParseMyID2  = Parse[MyID2]
	)

	r, err := ParseMyID("065sea8zyqr0dcpsk8e51g74zr")
	assert.Nil(t, err)
	r2, err := ParseMyID2("065s6jt2bvn68rxnnaezjv5wkg")
	assert.Nil(t, err)

	_, err = ParseMyID("065s6jt2bvn68rxnnaezjv5wkg")
	assert.Error(t, err, "wrong refid type")

	myr, err := NewMyID()
	assert.Nil(t, err)

	m := MyIDMatcher()
	assert.True(t, m.Match(r))
	assert.True(t, m.Match(&r))
	assert.True(t, m.Match(r.String()))
	assert.True(t, m.Match(r.Bytes()))
	assert.True(t, m.Match(myr))
	assert.True(t, m.Match(&myr))
	assert.True(t, m.Match(myr.String()))
	assert.True(t, m.Match(myr.Bytes()))
	assert.True(t, !m.Match(r2))
	assert.True(t, !m.Match(&r2))
	assert.True(t, !m.Match(r2.String()))
	assert.True(t, !m.Match(r2.Bytes()))
}
