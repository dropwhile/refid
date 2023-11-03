// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt55 = 55

type IDt55 struct {
	refid.ID
}

func (r *IDt55) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt55) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt55) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt55) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt55) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt55) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt55) toID() refid.ID {
	return r.ID
}

func (r IDt55) tagVal() byte {
	return tagIDt55
}

type NullIDt55 struct {
	refid.NullID
}

func (u *NullIDt55) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt55) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt55) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt55) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
