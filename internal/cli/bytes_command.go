// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright (c) 2024 Jan Hensel

package cli

import (
	"fmt"
	"os"

	"github.com/ja-he/count/internal/count"
)

// BytesCommand is a command for counting bytes.
type BytesCommand struct {
}

// Execute the bytes command.
func (command *BytesCommand) Execute(_ []string) error {
	count, err := count.BytesFromReader(os.Stdin)
	if err != nil {
		return fmt.Errorf("could not count bytes: %v", err)
	}
	fmt.Println(count)
	return nil
}
