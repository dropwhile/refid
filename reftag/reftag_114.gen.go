// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt114 = 114

type IDt114 struct {
	refid.ID
}

func (r *IDt114) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt114) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt114) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt114) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt114) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt114) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt114) toID() refid.ID {
	return r.ID
}

func (r IDt114) tagVal() byte {
	return tagIDt114
}

type NullIDt114 struct {
	refid.NullID
}

func (u *NullIDt114) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt114) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt114) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt114) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
