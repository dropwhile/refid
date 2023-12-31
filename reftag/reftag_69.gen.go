// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt69 = 69

type IDt69 struct {
	refid.ID
}

func (r *IDt69) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt69) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt69) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt69) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt69) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt69) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt69) toID() refid.ID {
	return r.ID
}

func (r IDt69) tagVal() byte {
	return tagIDt69
}

type NullIDt69 struct {
	refid.NullID
}

func (u *NullIDt69) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt69) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt69) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt69) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
