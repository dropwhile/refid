// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt246 = 246

type IDt246 struct {
	refid.ID
}

func (r *IDt246) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt246) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt246) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt246) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt246) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt246) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt246) toID() refid.ID {
	return r.ID
}

func (r IDt246) tagVal() byte {
	return tagIDt246
}

type NullIDt246 struct {
	refid.NullID
}

func (u *NullIDt246) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt246) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt246) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt246) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
