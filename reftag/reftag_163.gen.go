// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt163 = 163

type IDt163 struct {
	refid.ID
}

func (r *IDt163) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt163) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt163) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt163) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt163) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt163) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt163) toID() refid.ID {
	return r.ID
}

func (r IDt163) tagVal() byte {
	return tagIDt163
}

type NullIDt163 struct {
	refid.NullID
}

func (u *NullIDt163) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt163) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt163) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt163) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
