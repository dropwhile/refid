// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt48 = 48

type IDt48 struct {
	refid.ID
}

func (r *IDt48) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt48) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt48) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt48) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt48) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt48) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt48) toID() refid.ID {
	return r.ID
}

func (r IDt48) tagVal() byte {
	return tagIDt48
}

type NullIDt48 struct {
	refid.NullID
}

func (u *NullIDt48) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt48) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt48) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt48) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
