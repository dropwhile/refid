package refid

type Tagger byte

func NewTagger(tagVal byte) Tagger {
	return Tagger(tagVal)
}

func (t Tagger) New() (RefId, error) {
	return NewTagged(byte(t))
}

func (t Tagger) Parse(s string) (RefId, error) {
	return ParseTagged(byte(t), s)
}

func (t Tagger) HasTag(r RefId, tag byte) bool {
	return r.HasTag(tag)
}

func (t Tagger) IsTagged(r RefId) bool {
	return r.IsTagged()
}

func (t Tagger) AnyMatcher() AnyMatcher {
	return MatchAny(byte(t))
}

func (t Tagger) Tag() byte {
	return byte(t)
}

func (t Tagger) HasCorrectTag(r RefId) bool {
	return r.HasTag(byte(t))
}
