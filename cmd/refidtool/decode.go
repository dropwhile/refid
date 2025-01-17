// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/dropwhile/refid/v2"
)

type DecodeCmd struct {
	RefID string `arg:""`
}

func (cmd *DecodeCmd) Run() error {
	switch len(cmd.RefID) {
	case 26, 32, 22:
		// ok
	default:
		Fatal("invalid refid argument length")
	}

	xID, err := refid.Parse(cmd.RefID)
	if err != nil {
		Fatalf("error parsing refid: %s", err)
	}

	PrintRefID(xID)
	return nil
}
