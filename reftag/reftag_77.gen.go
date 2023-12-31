// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt77 = 77

type IDt77 struct {
	refid.ID
}

func (r *IDt77) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt77) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt77) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt77) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt77) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt77) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt77) toID() refid.ID {
	return r.ID
}

func (r IDt77) tagVal() byte {
	return tagIDt77
}

type NullIDt77 struct {
	refid.NullID
}

func (u *NullIDt77) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt77) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt77) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt77) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
