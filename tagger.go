package refid

type RefIdTagger byte

func NewRefIdTagger(tagVal byte) RefIdTagger {
	return RefIdTagger(tagVal)
}

func (rt RefIdTagger) New() (RefId, error) {
	return NewTagged(byte(rt))
}

func (rt RefIdTagger) MustNew() RefId {
	return MustNewTagged(byte(rt))
}

func (rt RefIdTagger) Parse(s string) (RefId, error) {
	return ParseTagged(byte(rt), s)
}

func (rt RefIdTagger) MustParse(s string) RefId {
	return MustParseTagged(byte(rt), s)
}

func (rt RefIdTagger) HasTag(r RefId, tag byte) bool {
	return r.HasTag(tag)
}

func (rt RefIdTagger) IsTagged(r RefId) bool {
	return r.IsTagged()
}

func (rt RefIdTagger) AnyMatcher() AnyMatcher {
	return MatchAny(byte(rt))
}

func (rt RefIdTagger) Tag() byte {
	return byte(rt)
}

func (rt RefIdTagger) HasCorrectTag(r RefId) bool {
	return r.HasTag(byte(rt))
}
