package refid

// A Tagger is a conveniece container for encoding and parsing
// [RefId]'s of a specific tag.
type Tagger byte

// NewTagger returns a new [Tagger] with tag
func NewTagger(tag byte) Tagger {
	return Tagger(tag)
}

// New generates a new RefId with tag set to the tag
// of the [Tagger]
func (t Tagger) New() (RefId, error) {
	return NewTagged(byte(t))
}

// Parse parses a [RefId], additionally enforcing that it is
// is tagged with the same tag as the [Tagger]
func (t Tagger) Parse(s string) (RefId, error) {
	return ParseTagged(byte(t), s)
}

// HasTag reports whether a [RefId] is tagged with
// the same tag as the [Tagger]
func (t Tagger) HasCorrectTag(r RefId) bool {
	return r.HasTag(byte(t))
}

// HasTag reports whether a [RefId] is tagged with a
// given tag
func (t Tagger) HasTag(r RefId, tag byte) bool {
	return r.HasTag(tag)
}

// IsTagged reports wheater a [RefId] is tagged at all.
// Note: This only checks that the [RefId] is tagged, not
// that it is tagged with the same tag as [Tagger]. For that
// functionality use [Tagger.HasCorrectTag].
func (t Tagger) IsTagged(r RefId) bool {
	return r.IsTagged()
}

// AnyMather returns an [AnyMatcher], which will
// match only against a [RefId] tagged with the same tag
// as the [Tagger]
func (t Tagger) AnyMatcher() AnyMatcher {
	return MatchAny(byte(t))
}

// Tag returns the tag of the [Tagger]
func (t Tagger) Tag() byte {
	return byte(t)
}
