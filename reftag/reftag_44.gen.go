// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt44 = 44

type IDt44 struct {
	refid.ID
}

func (r *IDt44) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt44) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt44) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt44) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt44) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt44) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt44) toID() refid.ID {
	return r.ID
}

func (r IDt44) tagVal() byte {
	return tagIDt44
}

type NullIDt44 struct {
	refid.NullID
}

func (u *NullIDt44) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt44) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt44) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt44) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}