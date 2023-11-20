// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt117 = 117

type IDt117 struct {
	refid.ID
}

func (r *IDt117) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt117) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt117) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt117) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt117) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt117) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt117) toID() refid.ID {
	return r.ID
}

func (r IDt117) tagVal() byte {
	return tagIDt117
}

type NullIDt117 struct {
	refid.NullID
}

func (u *NullIDt117) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt117) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt117) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt117) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}