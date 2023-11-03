// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt14 = 14

type IDt14 struct {
	refid.ID
}

func (r *IDt14) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt14) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt14) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt14) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt14) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt14) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt14) toID() refid.ID {
	return r.ID
}

func (r IDt14) tagVal() byte {
	return tagIDt14
}

type NullIDt14 struct {
	refid.NullID
}

func (u *NullIDt14) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt14) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt14) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt14) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
