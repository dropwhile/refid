// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt147 = 147

type IDt147 struct {
	refid.ID
}

func (r *IDt147) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt147) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt147) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt147) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt147) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt147) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt147) toID() refid.ID {
	return r.ID
}

func (r IDt147) tagVal() byte {
	return tagIDt147
}

type NullIDt147 struct {
	refid.NullID
}

func (u *NullIDt147) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt147) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt147) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt147) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
