// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt29 = 29

type IDt29 struct {
	refid.ID
}

func (r *IDt29) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt29) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt29) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt29) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt29) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt29) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt29) toID() refid.ID {
	return r.ID
}

func (r IDt29) tagVal() byte {
	return tagIDt29
}

type NullIDt29 struct {
	refid.NullID
}

func (u *NullIDt29) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt29) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt29) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt29) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
