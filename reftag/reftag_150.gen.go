// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt150 = 150

type IDt150 struct {
	refid.ID
}

func (r *IDt150) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt150) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt150) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt150) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt150) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt150) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt150) toID() refid.ID {
	return r.ID
}

func (r IDt150) tagVal() byte {
	return tagIDt150
}

type NullIDt150 struct {
	refid.NullID
}

func (u *NullIDt150) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt150) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt150) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt150) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
