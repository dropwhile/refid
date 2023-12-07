// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type verboseFlag bool

func (v verboseFlag) BeforeApply() error {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Debug().Msg("debug logging enabled")
	return nil
}

type CLI struct {
	// global options
	Verbose verboseFlag      `name:"verbose" short:"v" help:"enable verbose logging"`
	Version kong.VersionFlag `name:"version" short:"V" help:"Print version information and quit"`

	// subcommands
	Generate GenerateCmd `cmd:"" aliases:"gen" help:"Generate a new refid and print the details"`
	Decode   DecodeCmd   `cmd:"" aliases:"parse" help:"Decode and print the details of a refid"`
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stderr,
		PartsExclude: []string{zerolog.TimestampFieldName},
	})

	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("refidtool"),
		kong.Description("A tool for working with refids"),
		kong.UsageOnError(),
		kong.Vars{
			"version": "1.0.8",
		},
	)
	err := ctx.Run(&cli)
	ctx.FatalIfErrorf(err)
}
