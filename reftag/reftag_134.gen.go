// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt134 = 134

type IDt134 struct {
	refid.ID
}

func (r *IDt134) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt134) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt134) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt134) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt134) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt134) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt134) toID() refid.ID {
	return r.ID
}

func (r IDt134) tagVal() byte {
	return tagIDt134
}

type NullIDt134 struct {
	refid.NullID
}

func (u *NullIDt134) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt134) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt134) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt134) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
