// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt58 = 58

type IDt58 struct {
	refid.ID
}

func (r *IDt58) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt58) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt58) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt58) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt58) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt58) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt58) toID() refid.ID {
	return r.ID
}

func (r IDt58) tagVal() byte {
	return tagIDt58
}

type NullIDt58 struct {
	refid.NullID
}

func (u *NullIDt58) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt58) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt58) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt58) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
