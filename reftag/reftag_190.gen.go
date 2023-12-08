// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt190 = 190

type IDt190 struct {
	refid.ID
}

func (r *IDt190) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt190) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt190) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt190) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt190) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt190) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt190) toID() refid.ID {
	return r.ID
}

func (r IDt190) tagVal() byte {
	return tagIDt190
}

type NullIDt190 struct {
	refid.NullID
}

func (u *NullIDt190) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt190) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt190) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt190) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
