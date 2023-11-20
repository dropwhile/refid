// Code generated by refidgen. DO NOT EDIT.
// generated from: reftag.go

package reftag

import (
	"fmt"

	"github.com/dropwhile/refid"
)

const tagIDt182 = 182

type IDt182 struct {
	refid.ID
}

func (r *IDt182) Validate(err error) error {
	if err != nil {
		return err
	}
	if !r.ID.HasTag(tagIDt182) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (r *IDt182) Scan(src interface{}) error {
	err := r.ID.Scan(src)
	return r.Validate(err)
}

func (r *IDt182) UnmarshalJSON(b []byte) error {
	err := r.ID.UnmarshalJSON(b)
	return r.Validate(err)
}

func (r *IDt182) UnmarshalBinary(b []byte) error {
	err := r.ID.UnmarshalBinary(b)
	return r.Validate(err)
}

func (r *IDt182) UnmarshalText(b []byte) error {
	err := r.ID.UnmarshalText(b)
	return r.Validate(err)
}

func (r IDt182) toID() refid.ID {
	return r.ID
}

func (r IDt182) tagVal() byte {
	return tagIDt182
}

type NullIDt182 struct {
	refid.NullID
}

func (u *NullIDt182) Validate(err error) error {
	if err != nil {
		return err
	}
	n := u.NullID
	if n.Valid && !n.ID.HasTag(tagIDt182) {
		return fmt.Errorf("wrong refid type")
	}
	return nil
}

func (u *NullIDt182) Scan(src interface{}) error {
	err := u.NullID.Scan(src)
	return u.Validate(err)
}

func (u *NullIDt182) UnmarshalJSON(b []byte) error {
	err := u.NullID.UnmarshalJSON(b)
	return u.Validate(err)
}