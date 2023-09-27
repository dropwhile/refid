package refid

import "fmt"

type Requirement func(RefID) error

func HasTag(tag byte) Requirement {
	return func(r RefID) error {
		if !r.HasTag(tag) {
			return fmt.Errorf("RefID tag mismatch: go %d, expected %d", r[tagIndex], tag)
		}
		return nil
	}
}

func HasType(t Type) Requirement {
	return func(r RefID) error {
		if !r.HasType(t) {
			return fmt.Errorf("RefID type mismatch: got %s, expected %s", r.Type(), t)
		}
		return nil
	}
}
