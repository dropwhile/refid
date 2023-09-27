// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

// A Tagger is a conveniece container for encoding and parsing
// [RefID]'s of a specific tag.
type Tagger byte

// NewTagger returns a new [Tagger] with tag
func NewTagger(tag byte) Tagger {
	return Tagger(tag)
}

// New generates a new [TimePrefix] type [RefID] with tag set to the tag
// of the [Tagger]
func (t Tagger) New() (RefID, error) {
	return NewTagged(byte(t))
}

// NewRandom generates a new [RandomPrefix] type [RefID] with tag set to the tag
// of the [Tagger]
func (t Tagger) NewRandom() (RefID, error) {
	return NewRandomTagged(byte(t))
}

// Parse parses a [RefID], additionally enforcing that it is
// tagged with the same tag as the [Tagger]
func (t Tagger) Parse(s string) (RefID, error) {
	return ParseWithRequire(s, HasTag(byte(t)))
}

// ParseWithRequire parses a textual RefID representation (same formats as
// Parse), enforcing that it is tagged with the same tag as the [Tagger],
// while additionally requiring each reqs [Requirement] to pass, and returns
// a [RefID].
//
// Returns an error if RefID fails to parse, is not tagged with the same tag
// as [Tagger],  or if any of the reqs Requirements fail.
//
// Example:
//
//	ParseWithRequire("afd661f4f2tg2vr3dca92qp6k8", HasType(RandomPrefix))
func (t Tagger) ParseWithRequire(s string, reqs ...Requirement) (RefID, error) {
	reqs = append(reqs, HasTag(byte(t)))
	return ParseWithRequire(s, reqs...)
}

// HasTag reports whether a [RefID] is tagged with
// the same tag as the [Tagger]
func (t Tagger) HasCorrectTag(r RefID) bool {
	return r.HasTag(byte(t))
}

// HasTag reports whether a [RefID] is tagged with a
// given tag
func (t Tagger) HasTag(r RefID, tag byte) bool {
	return r.HasTag(tag)
}

// IsTagged reports wheater a [RefID] is tagged at all.
// Note: This only checks that the [RefID] is tagged, not
// that it is tagged with the same tag as [Tagger]. For that
// functionality use [Tagger.HasCorrectTag].
func (t Tagger) IsTagged(r RefID) bool {
	return r.IsTagged()
}

// AnyMather returns an [AnyMatcher], which will
// match only against a [RefID] tagged with the same tag
// as the [Tagger]
func (t Tagger) AnyMatcher() AnyMatcher {
	return MatchAny(byte(t))
}

// Tag returns the tag of the [Tagger]
func (t Tagger) Tag() byte {
	return byte(t)
}
