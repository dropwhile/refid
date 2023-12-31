// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt222 = 222

type IDt222 struct {
	refid.ID
}

func (r *IDt222) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt222) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt222) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt222) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt222) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt222) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt222) toID() refid.ID {
	return r.ID
}

func (r IDt222) tagVal() byte {
	return tagIDt222
}

type NullIDt222 struct {
	refid.NullID
}

func (u *NullIDt222) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt222) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt222) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt222) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
