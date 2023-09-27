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

// New generates a new RefID with tag set to the tag
// of the [Tagger]
func (t Tagger) New() (RefID, error) {
	return NewTagged(byte(t))
}

// Parse parses a [RefID], additionally enforcing that it is
// is tagged with the same tag as the [Tagger]
func (t Tagger) Parse(s string) (RefID, error) {
	return ParseTagged(byte(t), s)
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
