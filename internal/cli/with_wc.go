// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright (c) 2024 Jan Hensel

package cli

import (
	"fmt"
	"os"

	"github.com/ja-he/count/internal/wcwrap"
)

// WithWCCommand is the struct for the with-wc command.
type WithWCCommand struct {
}

// Execute the with-wc command.
func (c *WithWCCommand) Execute(_ []string) error {
	err := wcwrap.InvokeWithArgs(os.Stdin, os.Stdout)
	if err != nil {
		return fmt.Errorf("error invoking wc: %w", err)
	}

	return nil
}
