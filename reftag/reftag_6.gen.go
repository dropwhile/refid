// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt6 = 6

type IDt6 struct {
	refid.ID
}

func (r *IDt6) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt6) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt6) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt6) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt6) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt6) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt6) toID() refid.ID {
	return r.ID
}

func (r IDt6) tagVal() byte {
	return tagIDt6
}

type NullIDt6 struct {
	refid.NullID
}

func (u *NullIDt6) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt6) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt6) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt6) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
