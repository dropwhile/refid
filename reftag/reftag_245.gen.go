// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt245 = 245

type IDt245 struct {
	refid.ID
}

func (r *IDt245) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt245) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt245) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt245) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt245) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt245) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt245) toID() refid.ID {
	return r.ID
}

func (r IDt245) tagVal() byte {
	return tagIDt245
}

type NullIDt245 struct {
	refid.NullID
}

func (u *NullIDt245) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt245) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt245) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt245) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
