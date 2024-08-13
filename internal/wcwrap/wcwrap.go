// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright (c) 2024 Jan Hensel

// Package wcwrap is a wrapper over `wc`. Why would you write something like this package...?
package wcwrap

import (
	"fmt"
	"io"
	"os/exec"
)

// InvokeWithArgs invokes `wc` with the given arguments.
func InvokeWithArgs(stdin io.Reader, stdout io.Writer, args ...string) error {
	cmd := exec.Command("wc", args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("wc failed: %w", err)
	}

	return nil
}
