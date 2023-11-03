// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt72 = 72

type IDt72 struct {
	refid.ID
}

func (r *IDt72) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt72) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt72) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt72) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt72) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt72) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt72) toID() refid.ID {
	return r.ID
}

func (r IDt72) tagVal() byte {
	return tagIDt72
}

type NullIDt72 struct {
	refid.NullID
}

func (u *NullIDt72) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt72) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt72) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt72) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
