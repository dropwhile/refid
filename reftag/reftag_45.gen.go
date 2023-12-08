// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt45 = 45

type IDt45 struct {
	refid.ID
}

func (r *IDt45) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt45) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt45) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt45) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt45) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt45) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt45) toID() refid.ID {
	return r.ID
}

func (r IDt45) tagVal() byte {
	return tagIDt45
}

type NullIDt45 struct {
	refid.NullID
}

func (u *NullIDt45) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt45) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt45) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt45) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
