// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: Copyright (c) 2024 Jan Hensel

package count

import (
	"fmt"
	"io"
)

// Count is a count.
type Count int

// BytesFromReader reads from a reader and returns the number of bytes read.
func BytesFromReader(r io.Reader) (Count, error) {
	count := Count(0)

	buf := make([]byte, 4096)
	for {
		n, err := r.Read(buf)
		count += Count(n)
		if err != nil && err != io.EOF {
			return 0, fmt.Errorf("read: %w", err)
		}
		if err == io.EOF {
			break
		}
	}

	return Count(count), nil
}
