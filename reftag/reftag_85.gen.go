// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt85 = 85

type IDt85 struct {
	refid.ID
}

func (r *IDt85) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt85) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt85) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt85) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt85) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt85) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt85) toID() refid.ID {
	return r.ID
}

func (r IDt85) tagVal() byte {
	return tagIDt85
}

type NullIDt85 struct {
	refid.NullID
}

func (u *NullIDt85) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt85) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt85) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt85) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
