// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:     "reftool",
	Short:   "A tool for working with refids",
	Version: "0.0.1",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
			log.Debug().Msg("debug logging enabled")
		}
	},
}

func Execute() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stderr,
		PartsExclude: []string{zerolog.TimestampFieldName},
	})
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// execute
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
