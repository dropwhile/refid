// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt15 = 15

type IDt15 struct {
	refid.ID
}

func (r *IDt15) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt15) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt15) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt15) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt15) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt15) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt15) toID() refid.ID {
	return r.ID
}

func (r IDt15) tagVal() byte {
	return tagIDt15
}

type NullIDt15 struct {
	refid.NullID
}

func (u *NullIDt15) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt15) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt15) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt15) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
