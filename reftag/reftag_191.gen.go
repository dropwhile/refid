// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt191 = 191

type IDt191 struct {
	refid.ID
}

func (r *IDt191) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt191) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt191) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt191) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt191) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt191) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt191) toID() refid.ID {
	return r.ID
}

func (r IDt191) tagVal() byte {
	return tagIDt191
}

type NullIDt191 struct {
	refid.NullID
}

func (u *NullIDt191) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt191) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt191) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt191) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}