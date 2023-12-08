// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt152 = 152

type IDt152 struct {
	refid.ID
}

func (r *IDt152) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt152) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt152) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt152) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt152) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt152) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt152) toID() refid.ID {
	return r.ID
}

func (r IDt152) tagVal() byte {
	return tagIDt152
}

type NullIDt152 struct {
	refid.NullID
}

func (u *NullIDt152) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt152) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt152) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt152) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
