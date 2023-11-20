// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt230 = 230

type IDt230 struct {
	refid.ID
}

func (r *IDt230) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt230) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt230) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt230) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt230) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt230) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt230) toID() refid.ID {
	return r.ID
}

func (r IDt230) tagVal() byte {
	return tagIDt230
}

type NullIDt230 struct {
	refid.NullID
}

func (u *NullIDt230) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt230) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt230) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt230) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}