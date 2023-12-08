// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt105 = 105

type IDt105 struct {
	refid.ID
}

func (r *IDt105) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt105) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt105) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt105) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt105) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt105) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt105) toID() refid.ID {
	return r.ID
}

func (r IDt105) tagVal() byte {
	return tagIDt105
}

type NullIDt105 struct {
	refid.NullID
}

func (u *NullIDt105) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt105) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt105) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt105) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
