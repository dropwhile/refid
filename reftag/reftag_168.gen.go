// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt168 = 168

type IDt168 struct {
	refid.ID
}

func (r *IDt168) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt168) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt168) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt168) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt168) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt168) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt168) toID() refid.ID {
	return r.ID
}

func (r IDt168) tagVal() byte {
	return tagIDt168
}

type NullIDt168 struct {
	refid.NullID
}

func (u *NullIDt168) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt168) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt168) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt168) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}