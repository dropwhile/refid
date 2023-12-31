// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt95 = 95

type IDt95 struct {
	refid.ID
}

func (r *IDt95) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt95) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt95) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt95) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt95) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt95) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt95) toID() refid.ID {
	return r.ID
}

func (r IDt95) tagVal() byte {
	return tagIDt95
}

type NullIDt95 struct {
	refid.NullID
}

func (u *NullIDt95) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt95) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt95) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt95) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
