// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt109 = 109

type IDt109 struct {
	refid.ID
}

func (r *IDt109) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt109) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt109) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt109) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt109) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt109) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt109) toID() refid.ID {
	return r.ID
}

func (r IDt109) tagVal() byte {
	return tagIDt109
}

type NullIDt109 struct {
	refid.NullID
}

func (u *NullIDt109) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt109) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt109) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt109) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
