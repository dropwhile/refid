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
		refIdTxtLen := len(args[0])
		var parseFunc func(string) (refid.RefId, error)
		switch refIdTxtLen {
		case 0:
			log.Fatal().Msg("no refid argument provided")
		case 26: // native
			parseFunc = refid.Parse
		case 32: // hex
			parseFunc = refid.FromHexString
		case 22: // base64
			parseFunc = refid.FromBase64String
		default:
			log.Fatal().Msg("invalid refid argument length")
		}

		refId, err := parseFunc(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("invalid refid argument length")
		}

		tx := refId.Time()
		fmt.Printf("native enc:   %s\n", refId.String())
		fmt.Printf("hex enc:      %s\n", refId.ToHexString())
		fmt.Printf("base64 enc:   %s\n", refId.ToBase64String())
		fmt.Printf("tag value:    %d\n", refId.Tag())
		fmt.Printf("time(string): %s\n", tx.Format(time.RFC3339Nano))
		fmt.Printf("time(micros): %d\n", tx.UnixMicro())
	},
}
