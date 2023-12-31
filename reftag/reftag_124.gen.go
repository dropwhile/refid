// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt124 = 124

type IDt124 struct {
	refid.ID
}

func (r *IDt124) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt124) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt124) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt124) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt124) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt124) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt124) toID() refid.ID {
	return r.ID
}

func (r IDt124) tagVal() byte {
	return tagIDt124
}

type NullIDt124 struct {
	refid.NullID
}

func (u *NullIDt124) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt124) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt124) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt124) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
