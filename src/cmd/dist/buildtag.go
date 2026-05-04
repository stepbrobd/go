// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"go/build/constraint"
)

// matchexpr parses and evaluates the //go:build expression x.
func matchexpr(x string) (matched bool, _ error) {
	c, err := constraint.Parse("//go:build " + x)
	if err != nil {
		return false, fmt.Errorf("parsing //go:build line: %v", err)
	}
	return c.Eval(matchtag), nil
}
