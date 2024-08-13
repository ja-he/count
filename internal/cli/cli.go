// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright (c) 2024 Jan Hensel

// Package cli provides a CLI interface for the count command.
package cli

// Count is the CLI for the count binary.
type Count struct {
	Bytes  BytesCommand  `command:"bytes"`
	WrapWC WithWCCommand `command:"with-wc"`
}
