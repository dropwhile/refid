// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt136 = 136

type IDt136 struct {
	refid.ID
}

func (r *IDt136) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt136) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt136) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt136) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt136) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt136) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt136) toID() refid.ID {
	return r.ID
}

func (r IDt136) tagVal() byte {
	return tagIDt136
}

type NullIDt136 struct {
	refid.NullID
}

func (u *NullIDt136) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt136) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt136) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt136) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
