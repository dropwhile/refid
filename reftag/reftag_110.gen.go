// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid/v2"
)

const tagIDt110 = 110

type IDt110 struct {
	refid.ID
}

func (r *IDt110) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt110) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt110) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt110) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt110) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt110) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt110) toID() refid.ID {
	return r.ID
}

func (r IDt110) tagVal() byte {
	return tagIDt110
}

type NullIDt110 struct {
	refid.NullID
}

func (u *NullIDt110) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt110) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt110) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt110) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}
