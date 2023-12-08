// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt185 = 185

type IDt185 struct {
	refid.ID
}

func (r *IDt185) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt185) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt185) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt185) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt185) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt185) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt185) toID() refid.ID {
	return r.ID
}

func (r IDt185) tagVal() byte {
	return tagIDt185
}

type NullIDt185 struct {
	refid.NullID
}

func (u *NullIDt185) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt185) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt185) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt185) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
