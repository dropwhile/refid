// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt3 = 3

type IDt3 struct {
	refid.ID
}

func (r *IDt3) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt3) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt3) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt3) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt3) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt3) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt3) toID() refid.ID {
	return r.ID
}

func (r IDt3) tagVal() byte {
	return tagIDt3
}

type NullIDt3 struct {
	refid.NullID
}

func (u *NullIDt3) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt3) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt3) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt3) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}