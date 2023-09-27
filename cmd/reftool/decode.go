// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/dropwhile/refid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decodeCmd)
}

var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"parse"},
	Short:   "Decode and print the details of a refid",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args[0]) {
		case 26, 32, 22:
			// ok
		default:
			log.Fatal().Msg("invalid refid argument length")
		}

		xID, err := refid.Parse(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("error parsing refid")
		}

		tx := xID.Time().UTC()
		fmt.Printf("native enc:   %s\n", xID.String())
		fmt.Printf("hex enc:      %s\n", xID.ToHexString())
		fmt.Printf("base64 enc:   %s\n", xID.ToBase64String())
		fmt.Printf("tag value:    %d\n", xID.Tag())
		fmt.Printf("type value:   %s\n", xID.Type())
		fmt.Printf("time(string): %s\n", tx.Format(time.RFC3339Nano))
		fmt.Printf("time(micros): %d\n", tx.UnixMicro())
	},
}
