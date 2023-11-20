// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt16 = 16

type IDt16 struct {
	refid.ID
}

func (r *IDt16) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt16) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt16) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt16) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt16) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt16) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt16) toID() refid.ID {
	return r.ID
}

func (r IDt16) tagVal() byte {
	return tagIDt16
}

type NullIDt16 struct {
	refid.NullID
}

func (u *NullIDt16) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt16) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt16) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt16) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}