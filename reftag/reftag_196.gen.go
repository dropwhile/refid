// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt196 = 196

type IDt196 struct {
	refid.ID
}

func (r *IDt196) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt196) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt196) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt196) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt196) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt196) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt196) toID() refid.ID {
	return r.ID
}

func (r IDt196) tagVal() byte {
	return tagIDt196
}

type NullIDt196 struct {
	refid.NullID
}

func (u *NullIDt196) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt196) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt196) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt196) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
