package refid

type AnyMatcher struct {
	tag byte
}

func MatchAny(tag byte) AnyMatcher {
	return AnyMatcher{tag}
}

// Match satisfies sqlmock.Argument interface
func (a AnyMatcher) Match(v interface{}) bool {
	var refId RefId
	var err error
	switch x := v.(type) {
	case string:
		refId, err = Parse(x)
	case []byte:
		refId, err = FromBytes(x)
	case RefId:
		refId = x
	default:
		return false
	}
	if err != nil {
		return false
	}
	if a.tag != 0 {
		return refId.HasTag(a.tag)
	}
	return true
}
