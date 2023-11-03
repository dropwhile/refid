// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt175 = 175

type IDt175 struct {
	refid.ID
}

func (r *IDt175) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt175) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt175) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt175) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt175) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt175) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt175) toID() refid.ID {
	return r.ID
}

func (r IDt175) tagVal() byte {
	return tagIDt175
}

type NullIDt175 struct {
	refid.NullID
}

func (u *NullIDt175) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt175) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt175) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt175) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
