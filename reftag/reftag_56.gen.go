// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt56 = 56

type IDt56 struct {
	refid.ID
}

func (r *IDt56) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt56) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt56) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt56) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt56) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt56) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt56) toID() refid.ID {
	return r.ID
}

func (r IDt56) tagVal() byte {
	return tagIDt56
}

type NullIDt56 struct {
	refid.NullID
}

func (u *NullIDt56) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt56) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt56) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt56) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}