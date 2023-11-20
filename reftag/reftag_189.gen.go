// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt189 = 189

type IDt189 struct {
	refid.ID
}

func (r *IDt189) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt189) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt189) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt189) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt189) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt189) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt189) toID() refid.ID {
	return r.ID
}

func (r IDt189) tagVal() byte {
	return tagIDt189
}

type NullIDt189 struct {
	refid.NullID
}

func (u *NullIDt189) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt189) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt189) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt189) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}