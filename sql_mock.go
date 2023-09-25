package refid

type AnyMatcher struct {
	tag byte
}

func MatchAny(tag byte) AnyMatcher {
	return AnyMatcher{tag}
}

// Match satisfies sqlmock.Argument interface
func (a AnyMatcher) Match(v interface{}) bool {
	var r RefId
	var err error
	switch x := v.(type) {
	case string:
		r, err = Parse(x)
	case []byte:
		r, err = FromBytes(x)
	case RefId:
		r = x
	default:
		return false
	}
	if err != nil {
		return false
	}
	if a.tag != 0 {
		return r.HasTag(a.tag)
	}
	return true
}
