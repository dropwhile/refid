// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt200 = 200

type IDt200 struct {
	refid.ID
}

func (r *IDt200) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt200) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt200) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt200) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt200) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt200) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt200) toID() refid.ID {
	return r.ID
}

func (r IDt200) tagVal() byte {
	return tagIDt200
}

type NullIDt200 struct {
	refid.NullID
}

func (u *NullIDt200) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt200) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt200) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt200) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
