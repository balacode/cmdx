// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-24 03:05:23 5EE2CC                          cmdx/[module_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in module.go use:
    go test --run Test_mdle_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_mdle_
func Test_mdle_(t *testing.T) {
	zr.TBegin(t)
	// module
} //                                                                  Test_mdle_

//end
