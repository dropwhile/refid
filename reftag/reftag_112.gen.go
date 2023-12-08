// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt112 = 112

type IDt112 struct {
	refid.ID
}

func (r *IDt112) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt112) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt112) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt112) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt112) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt112) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt112) toID() refid.ID {
	return r.ID
}

func (r IDt112) tagVal() byte {
	return tagIDt112
}

type NullIDt112 struct {
	refid.NullID
}

func (u *NullIDt112) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt112) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt112) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt112) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
