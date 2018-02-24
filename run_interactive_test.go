// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:47 8DDEA8                 [cmdx/run_interactive_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in run_interactive.go use:
    go test --run Test_runi_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zirconium

// go test --run Test_runi_runInteractive_
func Test_runi_runInteractive_(t *testing.T) {
	zr.TBegin(t)
	// runInteractive(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		runInteractive(cmd, args)
	}
	test(Command{}, []string{})
} //                                                   Test_runi_runInteractive_

//end
