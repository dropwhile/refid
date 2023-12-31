// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt66 = 66

type IDt66 struct {
	refid.ID
}

func (r *IDt66) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt66) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt66) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt66) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt66) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt66) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt66) toID() refid.ID {
	return r.ID
}

func (r IDt66) tagVal() byte {
	return tagIDt66
}

type NullIDt66 struct {
	refid.NullID
}

func (u *NullIDt66) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt66) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt66) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt66) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
