// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt63 = 63

type IDt63 struct {
	refid.ID
}

func (r *IDt63) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt63) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt63) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt63) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt63) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt63) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt63) toID() refid.ID {
	return r.ID
}

func (r IDt63) tagVal() byte {
	return tagIDt63
}

type NullIDt63 struct {
	refid.NullID
}

func (u *NullIDt63) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt63) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt63) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt63) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}