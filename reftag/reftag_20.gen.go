// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt20 = 20

type IDt20 struct {
	refid.ID
}

func (r *IDt20) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt20) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt20) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt20) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt20) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt20) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt20) toID() refid.ID {
	return r.ID
}

func (r IDt20) tagVal() byte {
	return tagIDt20
}

type NullIDt20 struct {
	refid.NullID
}

func (u *NullIDt20) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt20) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt20) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt20) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
