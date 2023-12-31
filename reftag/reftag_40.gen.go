// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt40 = 40

type IDt40 struct {
	refid.ID
}

func (r *IDt40) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt40) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt40) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt40) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt40) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt40) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt40) toID() refid.ID {
	return r.ID
}

func (r IDt40) tagVal() byte {
	return tagIDt40
}

type NullIDt40 struct {
	refid.NullID
}

func (u *NullIDt40) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt40) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt40) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt40) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
