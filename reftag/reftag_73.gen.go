// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt73 = 73

type IDt73 struct {
	refid.ID
}

func (r *IDt73) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt73) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt73) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt73) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt73) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt73) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt73) toID() refid.ID {
	return r.ID
}

func (r IDt73) tagVal() byte {
	return tagIDt73
}

type NullIDt73 struct {
	refid.NullID
}

func (u *NullIDt73) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt73) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt73) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt73) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
