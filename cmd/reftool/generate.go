package main

import (
	"fmt"
	"time"

	"github.com/dropwhile/refid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	TagValue uint8
	TimeAt   string
	Only     string
)

func init() {
	generateCmd.Flags().Uint8VarP(&TagValue, "tag-value", "t", 0, "tag value")
	generateCmd.Flags().StringVarP(&Only, "only", "o", "", "output only encoding result. optional argument: native, base64, hex")
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
		var refId refid.RefId
		if TagValue != 0 {
			refId = refid.Must(refid.NewTagged(TagValue))
		} else {
			refId = refid.Must(refid.New())
		}

		var ts time.Time
		if TimeAt != "" {
			var err error
			ts, err = time.Parse(time.RFC3339, TimeAt)
			if err != nil {
				log.Fatal().Err(err).Msg("error parsing datetime")
			}
			refId.SetTime(ts)
		}

		switch Only {
		case "base64":
			fmt.Println(refId.ToBase64String())
			return
		case "hex":
			fmt.Println(refId.ToHexString())
			return
		case "":
			tx := refId.Time()
			fmt.Printf("native enc:   %s\n", refId.String())
			fmt.Printf("hex enc:      %s\n", refId.ToHexString())
			fmt.Printf("base64 enc:   %s\n", refId.ToBase64String())
			fmt.Printf("tag value:    %d\n", refId.Tag())
			fmt.Printf("time(string): %s\n", tx.Format(time.RFC3339Nano))
			fmt.Printf("time(micros): %d\n", tx.UnixMicro())
		default:
			fmt.Println(refId.String())
			return
		}
	},
}
