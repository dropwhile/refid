// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt165 = 165

type IDt165 struct {
	refid.ID
}

func (r *IDt165) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt165) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt165) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt165) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt165) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt165) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt165) toID() refid.ID {
	return r.ID
}

func (r IDt165) tagVal() byte {
	return tagIDt165
}

type NullIDt165 struct {
	refid.NullID
}

func (u *NullIDt165) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt165) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt165) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt165) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
