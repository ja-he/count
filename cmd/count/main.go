// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright (c) 2024 Jan Hensel

// Package main is a binary for `count` which behaves very roughly like `wc`.
package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/ja-he/count/internal/cli"
)

func main() {
	var cliOpts cli.Count
	parser := flags.NewParser(&cliOpts, flags.Default)
	parser.SubcommandsOptional = false
	if _, err := parser.Parse(); err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "could not parse command line options: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
