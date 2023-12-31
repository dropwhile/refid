// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt139 = 139

type IDt139 struct {
	refid.ID
}

func (r *IDt139) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt139) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt139) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt139) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt139) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt139) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt139) toID() refid.ID {
	return r.ID
}

func (r IDt139) tagVal() byte {
	return tagIDt139
}

type NullIDt139 struct {
	refid.NullID
}

func (u *NullIDt139) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt139) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt139) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt139) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
