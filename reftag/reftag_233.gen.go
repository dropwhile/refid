// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt233 = 233

type IDt233 struct {
	refid.ID
}

func (r *IDt233) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt233) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt233) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt233) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt233) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt233) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt233) toID() refid.ID {
	return r.ID
}

func (r IDt233) tagVal() byte {
	return tagIDt233
}

type NullIDt233 struct {
	refid.NullID
}

func (u *NullIDt233) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt233) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt233) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt233) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
