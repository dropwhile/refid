// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt220 = 220

type IDt220 struct {
	refid.ID
}

func (r *IDt220) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt220) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt220) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt220) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt220) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt220) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt220) toID() refid.ID {
	return r.ID
}

func (r IDt220) tagVal() byte {
	return tagIDt220
}

type NullIDt220 struct {
	refid.NullID
}

func (u *NullIDt220) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt220) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt220) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt220) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
