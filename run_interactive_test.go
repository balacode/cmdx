// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-12 16:56:18 E9D31A                 cmdx/[run_interactive_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in run_interactive.go use:
//      go test --run Test_runi_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_runi_runInteractive_
func Test_runi_runInteractive_(t *testing.T) {
	zr.TBegin(t)
	// runInteractive(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		runInteractive(cmd, args)
	}
	test(Command{}, []string{})
} //                                                   Test_runi_runInteractive_

//end
