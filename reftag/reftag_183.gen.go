// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt183 = 183

type IDt183 struct {
	refid.ID
}

func (r *IDt183) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt183) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt183) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt183) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt183) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt183) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt183) toID() refid.ID {
	return r.ID
}

func (r IDt183) tagVal() byte {
	return tagIDt183
}

type NullIDt183 struct {
	refid.NullID
}

func (u *NullIDt183) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt183) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt183) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt183) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
