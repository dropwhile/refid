// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/dropwhile/refid/v2"
)

type GenerateCmd struct {
	TagValue uint8      `name:"tag-value" short:"t" help:"tag value"`
	Only     string     `name:"only" short:"o" enum:"native,base64,hex,all" default:"all" help:"output only encoding result. one of: ${enum}"`
	Random   bool       `name:"random" short:"r" help:"generate a RandomPrefixed instead of a TimePrefixed ref.ID"`
	When     *time.Time `name:"when" short:"w" help:"the date/time to use in the token, truncated to seconds resolution. Uses RFC3339 format"`
}

func (cmd *GenerateCmd) Run() error {
	var xID refid.ID
	if cmd.Random {
		xID = refid.Must(refid.NewRandom())
	} else {
		xID = refid.Must(refid.New())
	}

	if cmd.TagValue != 0 {
		xID.SetTag(cmd.TagValue)
	}

	if cmd.When != nil {
		_ = xID.SetTime(*cmd.When)
	}

	switch cmd.Only {
	case "base64":
		fmt.Println(xID.ToBase64String())
	case "hex":
		fmt.Println(xID.ToHexString())
	case "native":
		fmt.Println(xID.String())
	default:
		PrintRefID(xID)
	}
	return nil
}
