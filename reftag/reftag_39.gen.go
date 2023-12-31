// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt39 = 39

type IDt39 struct {
	refid.ID
}

func (r *IDt39) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt39) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt39) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt39) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt39) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt39) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt39) toID() refid.ID {
	return r.ID
}

func (r IDt39) tagVal() byte {
	return tagIDt39
}

type NullIDt39 struct {
	refid.NullID
}

func (u *NullIDt39) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt39) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt39) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt39) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
