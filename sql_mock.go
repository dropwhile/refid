package refid

// A matcher that supports the following interfaces:
//   - [github.com/pashagolub/pgxmock.Argument]
//   - [github.com/DATA-DOG/go-sqlmock.Argument]
type AnyMatcher struct {
	tag byte
}

// Create a [AnyMatcher] matcher that matches that matches against a specific
// Tag. Any valid RefIds that do not match the tag specified, will be considered
// not matching.
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
