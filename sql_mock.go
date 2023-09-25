package refid

// A matcher that supports the pgxmock.Argument interface
type AnyMatcher struct {
	tag byte
}

// Create a AnyMatcher matcher that supports the pgxmock.Argument interface,
// and matches against a specific Tag. Any valid RefIds that do not
// match the tag specified, will be considered not matching.
//
// If tag is 0, will support matching any RefId (tag is then ignored)
//
// Example usage:
//
//	mock.ExpectQuery("^INSERT INTO some_table (.+)").
//	 WithArgs(refid.MatchAny(1), 1).
//	 WillReturnRows(rows)
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
