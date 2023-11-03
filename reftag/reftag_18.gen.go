// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt18 = 18

type IDt18 struct {
	refid.ID
}

func (r *IDt18) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt18) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt18) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt18) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt18) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt18) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt18) toID() refid.ID {
	return r.ID
}

func (r IDt18) tagVal() byte {
	return tagIDt18
}

type NullIDt18 struct {
	refid.NullID
}

func (u *NullIDt18) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt18) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt18) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt18) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
