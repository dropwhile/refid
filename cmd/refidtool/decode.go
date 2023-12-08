// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/dropwhile/refid/v2"
	"github.com/rs/zerolog/log"
)

type DecodeCmd struct {
	RefID string `arg:""`
}

func (cmd *DecodeCmd) Run() error {
	switch len(cmd.RefID) {
	case 26, 32, 22:
		// ok
	default:
		log.Fatal().Msg("invalid refid argument length")
	}

	xID, err := refid.Parse(cmd.RefID)
	if err != nil {
		log.Fatal().Err(err).Msg("error parsing refid")
	}

	PrintRefID(xID)
	return nil
}
