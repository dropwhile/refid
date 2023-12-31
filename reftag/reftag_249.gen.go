// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt249 = 249

type IDt249 struct {
	refid.ID
}

func (r *IDt249) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt249) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt249) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt249) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt249) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt249) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt249) toID() refid.ID {
	return r.ID
}

func (r IDt249) tagVal() byte {
	return tagIDt249
}

type NullIDt249 struct {
	refid.NullID
}

func (u *NullIDt249) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt249) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt249) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt249) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
