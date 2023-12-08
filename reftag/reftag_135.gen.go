// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt135 = 135

type IDt135 struct {
	refid.ID
}

func (r *IDt135) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt135) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt135) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt135) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt135) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt135) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt135) toID() refid.ID {
	return r.ID
}

func (r IDt135) tagVal() byte {
	return tagIDt135
}

type NullIDt135 struct {
	refid.NullID
}

func (u *NullIDt135) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt135) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt135) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt135) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
