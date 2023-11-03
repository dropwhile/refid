// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt180 = 180

type IDt180 struct {
	refid.ID
}

func (r *IDt180) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt180) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt180) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt180) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt180) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt180) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt180) toID() refid.ID {
	return r.ID
}

func (r IDt180) tagVal() byte {
	return tagIDt180
}

type NullIDt180 struct {
	refid.NullID
}

func (u *NullIDt180) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt180) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt180) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt180) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
