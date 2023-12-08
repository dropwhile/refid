package reftag

import (
	"database/sql/driver"

	"github.com/dropwhile/refid/v2"
)

//go:generate go run ../tool/refidgen.go

type RefTagger[C any] interface {
	Scan(interface{}) error
	UnmarshalText([]byte) error
	UnmarshalBinary([]byte) error
	Validate(error) error
	Value() (driver.Value, error)
	toID() refid.ID
	tagVal() byte
	*C // non-interface type constraint element
}

func New[V any, T RefTagger[V]]() (V, error) {
	var ptr V
	r := T(&ptr)
	tv := r.tagVal()
	rid, err := refid.NewTagged(tv)
	if err != nil {
		return *r, err
	}
	err = r.UnmarshalBinary(rid.Bytes())
	err = r.Validate(err)
	return *r, err
}

func NewRandom[V any, T RefTagger[V]]() (V, error) {
	var ptr V
	r := T(&ptr)
	tv := r.tagVal()
	rid, err := refid.NewRandomTagged(tv)
	if err != nil {
		return *r, err
	}
	err = r.UnmarshalBinary(rid.Bytes())
	err = r.Validate(err)
	return *r, err
}

func Parse[V any, T RefTagger[V]](s string) (V, error) {
	var ptr V
	r := T(&ptr)
	err := r.UnmarshalText([]byte(s))
	err = r.Validate(err)
	return *r, err
}

func ParseWithRequire[V any, T RefTagger[V]](s string, reqs ...refid.Requirement) (V, error) {
	r, err := Parse[V, T](s)
	if err != nil {
		return r, err
	}

	for _, f := range reqs {
		t := T(&r)
		err = f(t.toID())
		if err != nil {
			return r, err
		}
	}
	return r, nil
}

func FromBytes[V any, T RefTagger[V]](input []byte) (V, error) {
	var ptr V
	r := T(&ptr)
	err := r.UnmarshalBinary(input)
	err = r.Validate(err)
	return *r, err
}

type Matcher[V any, T RefTagger[V]] struct{}

func (a Matcher[V, T]) Match(v interface{}) bool {
	var ptr V
	r := T(&ptr)
	var err error
	switch x := v.(type) {
	case T:
		r = x
	case V:
		r = T(&x)
	case refid.ID:
		err = r.UnmarshalBinary(x.Bytes())
	case *refid.ID:
		err = r.UnmarshalBinary(x.Bytes())
	case string:
		ptr, err = Parse[V, T](x)
		r = T(&ptr)
	case []byte:
		ptr, err = FromBytes[V, T](x)
		r = T(&ptr)
	default:
		return false
	}
	return r.Validate(err) == nil
}

func NewMatcher[V any, T RefTagger[V]]() Matcher[V, T] {
	return Matcher[V, T]{}
}
