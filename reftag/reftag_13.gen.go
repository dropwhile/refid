// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt13 = 13

type IDt13 struct {
	refid.ID
}

func (r *IDt13) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt13) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt13) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt13) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt13) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt13) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt13) toID() refid.ID {
	return r.ID
}

func (r IDt13) tagVal() byte {
	return tagIDt13
}

type NullIDt13 struct {
	refid.NullID
}

func (u *NullIDt13) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt13) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt13) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt13) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
