// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt120 = 120

type IDt120 struct {
	refid.ID
}

func (r *IDt120) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt120) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt120) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt120) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt120) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt120) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt120) toID() refid.ID {
	return r.ID
}

func (r IDt120) tagVal() byte {
	return tagIDt120
}

type NullIDt120 struct {
	refid.NullID
}

func (u *NullIDt120) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt120) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt120) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt120) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
