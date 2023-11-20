// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt215 = 215

type IDt215 struct {
	refid.ID
}

func (r *IDt215) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt215) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt215) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt215) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt215) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt215) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt215) toID() refid.ID {
	return r.ID
}

func (r IDt215) tagVal() byte {
	return tagIDt215
}

type NullIDt215 struct {
	refid.NullID
}

func (u *NullIDt215) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt215) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt215) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt215) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}