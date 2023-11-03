// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt169 = 169

type IDt169 struct {
	refid.ID
}

func (r *IDt169) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt169) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt169) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt169) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt169) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt169) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt169) toID() refid.ID {
	return r.ID
}

func (r IDt169) tagVal() byte {
	return tagIDt169
}

type NullIDt169 struct {
	refid.NullID
}

func (u *NullIDt169) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt169) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt169) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt169) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
