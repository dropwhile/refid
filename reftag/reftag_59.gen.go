// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt59 = 59

type IDt59 struct {
	refid.ID
}

func (r *IDt59) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt59) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt59) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt59) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt59) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt59) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt59) toID() refid.ID {
	return r.ID
}

func (r IDt59) tagVal() byte {
	return tagIDt59
}

type NullIDt59 struct {
	refid.NullID
}

func (u *NullIDt59) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt59) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt59) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt59) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}