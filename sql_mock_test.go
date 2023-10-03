// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

import (
	"testing"
)

func TestAnyMatcher_Match(t *testing.T) {
	tests := []struct {
		name string
		tag  byte
		arg  interface{}
		want bool
	}{
		{"refid: test zero matcher with zero tag", 0x00, Must(NewTagged(0x00)), true},
		{"refid: test zero matcher with nonzero tag", 0x00, Must(NewTagged(0x01)), true},
		{"refid: test tag:0x01 matcher with zero tag", 0x01, Must(NewTagged(0x00)), false},
		{"refid: test tag:0x01 matcher with 0x01 tag", 0x01, Must(NewTagged(0x01)), true},
		{"refid: test tag:0xff matcher with 0x01 tag", 0xff, Must(NewTagged(0x01)), false},
		{"refid: test tag:0xff matcher with 0xff tag", 0xff, Must(NewTagged(0xff)), true},
		{"refid: test tag:0xff matcher with no tag", 0xff, Must(New()), false},
		{"refid: test tag:0x00 matcher with no tag", 0xff, Must(New()), false},

		{"string: test zero matcher with zero value", 0x00, Must(NewTagged(0x00)).String(), true},
		{"string: test zero matcher with nonzero value", 0x00, Must(NewTagged(0x01)).String(), true},
		{"string: test tag:0x01 matcher with zero value", 0x01, Must(NewTagged(0x00)).String(), false},
		{"string: test tag:0x01 matcher with 0x01 value", 0x01, Must(NewTagged(0x01)).String(), true},
		{"string: test tag:0xff matcher with 0x01 value", 0xff, Must(NewTagged(0x01)).String(), false},
		{"string: test tag:0xff matcher with 0xff value", 0xff, Must(NewTagged(0xff)).String(), true},
		{"string: test tag:0xff matcher with no tag", 0xff, Must(New()).String(), false},
		{"string: test tag:0x00 matcher with no tag", 0xff, Must(New()).String(), false},

		{"bytes: test zero matcher with zero value", 0x00, Must(NewTagged(0x00)).Bytes(), true},
		{"bytes: test zero matcher with nonzero value", 0x00, Must(NewTagged(0x01)).Bytes(), true},
		{"bytes: test tag:0x01 matcher with zero value", 0x01, Must(NewTagged(0x00)).Bytes(), false},
		{"bytes: test tag:0x01 matcher with 0x01 value", 0x01, Must(NewTagged(0x01)).Bytes(), true},
		{"bytes: test tag:0xff matcher with 0x01 value", 0xff, Must(NewTagged(0x01)).Bytes(), false},
		{"bytes: test tag:0xff matcher with 0xff value", 0xff, Must(NewTagged(0xff)).Bytes(), true},
		{"bytes: test tag:0xff matcher with no tag", 0xff, Must(New()).Bytes(), false},
		{"bytes: test tag:0x00 matcher with no tag", 0xff, Must(New()).Bytes(), false},

		{"bad-refid: test zero matcher with zero value", 0x00, "nope", false},
		{"bad-refid: test zero matcher with nonzero value", 0x00, "nope", false},
		{"bad-refid: test tag:0x01 matcher with zero value", 0x01, "nope", false},
		{"bad-refid: test tag:0x01 matcher with 0x01 value", 0x01, "nope", false},
		{"bad-refid: test tag:0xff matcher with 0x01 value", 0xff, "nope", false},
		{"bad-refid: test tag:0xff matcher with 0xff value", 0xff, "nope", false},
		{"bad-refid: test tag:0xff matcher with no tag", 0xff, "nope", false},
		{"bad-refid: test tag:0x00 matcher with no tag", 0xff, "nope", false},

		{"invalid-gotype: test zero matcher with zero value", 0x00, int64(7), false},
		{"invalid-gotype: test zero matcher with nonzero value", 0x00, int64(7), false},
		{"invalid-gotype: test tag:0x01 matcher with zero value", 0x01, int64(7), false},
		{"invalid-gotype: test tag:0x01 matcher with 0x01 value", 0x01, int64(7), false},
		{"invalid-gotype: test tag:0xff matcher with 0x01 value", 0xff, int64(7), false},
		{"invalid-gotype: test tag:0xff matcher with 0xff value", 0xff, int64(7), false},
		{"invalid-gotype: test tag:0xff matcher with no tag", 0xff, int64(7), false},
		{"invalid-gotype: test tag:0x00 matcher with no tag", 0xff, int64(7), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AnyMatcher{
				tag: tt.tag,
			}
			if got := a.Match(tt.arg); got != tt.want {
				t.Errorf("AnyMatcher.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
