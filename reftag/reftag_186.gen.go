// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt186 = 186

type IDt186 struct {
	refid.ID
}

func (r *IDt186) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt186) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt186) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt186) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt186) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt186) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt186) toID() refid.ID {
	return r.ID
}

func (r IDt186) tagVal() byte {
	return tagIDt186
}

type NullIDt186 struct {
	refid.NullID
}

func (u *NullIDt186) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt186) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt186) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt186) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
