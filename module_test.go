// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-12 16:56:18 F75D24                          cmdx/[module_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in module.go use:
//      go test --run Test_mdle_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

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
