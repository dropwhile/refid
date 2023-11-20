// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt23 = 23

type IDt23 struct {
	refid.ID
}

func (r *IDt23) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt23) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt23) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt23) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt23) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt23) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt23) toID() refid.ID {
	return r.ID
}

func (r IDt23) tagVal() byte {
	return tagIDt23
}

type NullIDt23 struct {
	refid.NullID
}

func (u *NullIDt23) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt23) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt23) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt23) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}