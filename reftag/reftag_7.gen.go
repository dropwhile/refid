// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt7 = 7

type IDt7 struct {
	refid.ID
}

func (r *IDt7) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt7) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt7) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt7) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt7) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt7) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt7) toID() refid.ID {
	return r.ID
}

func (r IDt7) tagVal() byte {
	return tagIDt7
}

type NullIDt7 struct {
	refid.NullID
}

func (u *NullIDt7) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt7) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt7) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt7) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
