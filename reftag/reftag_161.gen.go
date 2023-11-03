// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt161 = 161

type IDt161 struct {
	refid.ID
}

func (r *IDt161) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt161) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt161) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt161) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt161) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt161) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt161) toID() refid.ID {
	return r.ID
}

func (r IDt161) tagVal() byte {
	return tagIDt161
}

type NullIDt161 struct {
	refid.NullID
}

func (u *NullIDt161) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt161) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt161) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt161) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
