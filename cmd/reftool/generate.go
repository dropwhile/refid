// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/dropwhile/refid"
	"github.com/spf13/cobra"
)

var (
	TagValue uint8
	TimeAt   string
	Only     string
	Random   bool
)

func init() {
	generateCmd.Flags().Uint8VarP(&TagValue, "tag-value", "t", 0, "tag value")
	generateCmd.Flags().StringVarP(&Only, "only", "o", "", "output only encoding result. optional argument: native, base64, hex")
	generateCmd.Flags().BoolVarP(&Random, "random", "r", false, "generate a RefIdRandom instead of a standard type")
	generateCmd.Flags().Lookup("only").NoOptDefVal = "native"
	generateCmd.Flags().StringVarP(
		&TimeAt, "when", "w", "",
		"the date/time to use in the token, truncated to seconds resolution. Uses RFC3339 format",
	)
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate a new refid and print the details",
	Run: func(cmd *cobra.Command, args []string) {
		var xID refid.RefID
		if TagValue != 0 {
			if Random {
				xID = refid.Must(refid.NewRandomTagged(TagValue))
			} else {
				xID = refid.Must(refid.NewTagged(TagValue))
			}
		} else {
			if Random {
				xID = refid.Must(refid.NewRandom())
			} else {
				xID = refid.Must(refid.New())
			}
		}

		switch Only {
		case "base64":
			fmt.Println(xID.ToBase64String())
		case "hex":
			fmt.Println(xID.ToHexString())
		case "":
			tx := xID.Time().UTC()
			fmt.Printf("native enc:   %s\n", xID.String())
			fmt.Printf("hex enc:      %s\n", xID.ToHexString())
			fmt.Printf("base64 enc:   %s\n", xID.ToBase64String())
			fmt.Printf("tag value:    %d\n", xID.Tag())
			fmt.Printf("type:         %s\n", xID.Type())
			fmt.Printf("time(string): %s\n", tx.Format(time.RFC3339Nano))
			fmt.Printf("time(millis): %d\n", tx.UnixMilli())
		default:
			fmt.Println(xID.String())
		}
	},
}
