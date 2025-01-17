// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"log/slog"
	"runtime/debug"

	"github.com/alecthomas/kong"
)

var Version string

func GetVersion() string {
	if len(Version) > 0 {
		return Version
	}

	if bi, ok := debug.ReadBuildInfo(); ok {
		// If no main version is available, Go defaults it to (devel)
		if bi.Main.Version != "(devel)" {
			return bi.Main.Version
		}
	}

	return "unknown"
}

type CLI struct {
	// global options
	Verbose bool             `name:"verbose" short:"v" help:"enable verbose logging"`
	Version kong.VersionFlag `name:"version" short:"V" help:"Print version information and quit"`

	// subcommands
	Generate GenerateCmd `cmd:"" aliases:"gen" help:"Generate a new refid and print the details"`
	Decode   DecodeCmd   `cmd:"" aliases:"parse" help:"Decode and print the details of a refid"`
}

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("refidtool"),
		kong.Description("A tool for working with refids"),
		kong.UsageOnError(),
		kong.Vars{
			"version": GetVersion(),
		},
	)
	if cli.Verbose {
		logLevel.Set(slog.LevelDebug)
	}
	err := ctx.Run(&cli)
	ctx.FatalIfErrorf(err)
}
