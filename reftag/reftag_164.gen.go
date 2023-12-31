// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt164 = 164

type IDt164 struct {
	refid.ID
}

func (r *IDt164) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt164) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt164) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt164) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt164) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt164) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt164) toID() refid.ID {
	return r.ID
}

func (r IDt164) tagVal() byte {
	return tagIDt164
}

type NullIDt164 struct {
	refid.NullID
}

func (u *NullIDt164) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt164) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt164) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt164) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
