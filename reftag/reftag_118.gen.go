// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt118 = 118

type IDt118 struct {
	refid.ID
}

func (r *IDt118) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt118) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt118) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt118) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt118) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt118) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt118) toID() refid.ID {
	return r.ID
}

func (r IDt118) tagVal() byte {
	return tagIDt118
}

type NullIDt118 struct {
	refid.NullID
}

func (u *NullIDt118) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt118) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt118) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt118) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
