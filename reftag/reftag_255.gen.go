// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt255 = 255

type IDt255 struct {
	refid.ID
}

func (r *IDt255) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt255) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt255) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt255) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt255) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt255) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt255) toID() refid.ID {
	return r.ID
}

func (r IDt255) tagVal() byte {
	return tagIDt255
}

type NullIDt255 struct {
	refid.NullID
}

func (u *NullIDt255) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt255) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt255) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt255) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
