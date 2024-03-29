// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                      cmdx/[main_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

//  to test all items in main.go use:
//      go test --run Test_main_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Main Function

// go test --run Test_main_
func Test_main_(t *testing.T) {
	zr.TBegin(t)
	// main()
	//
	test := func() {
		// TODO: main()
	}
	test()
}

// end
