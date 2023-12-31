// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt101 = 101

type IDt101 struct {
	refid.ID
}

func (r *IDt101) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt101) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt101) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt101) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt101) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt101) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt101) toID() refid.ID {
	return r.ID
}

func (r IDt101) tagVal() byte {
	return tagIDt101
}

type NullIDt101 struct {
	refid.NullID
}

func (u *NullIDt101) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt101) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt101) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt101) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
