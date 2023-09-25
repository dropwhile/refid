package refid

type RefIdTagger byte

func NewRefIdTagger(tagVal byte) RefIdTagger {
	return RefIdTagger(tagVal)
}

func (r RefIdTagger) New() (RefId, error) {
	return NewTagged(byte(r))
}

func (r RefIdTagger) MustNew() RefId {
	return MustNewTagged(byte(r))
}

func (r RefIdTagger) Parse(s string) (RefId, error) {
	return ParseTagged(byte(r), s)
}

func (r RefIdTagger) MustParse(s string) RefId {
	return MustParseTagged(byte(r), s)
}

func (r RefIdTagger) HasTag(refId RefId, tag byte) bool {
	return refId.HasTag(tag)
}

func (r RefIdTagger) IsTagged(refId RefId) bool {
	return refId.IsTagged()
}

func (r RefIdTagger) AnyMatcher() AnyMatcher {
	return MatchAny(byte(r))
}

func (r RefIdTagger) Tag() byte {
	return byte(r)
}

func (r RefIdTagger) HasCorrectTag(refId RefId) bool {
	return refId.HasTag(byte(r))
}
