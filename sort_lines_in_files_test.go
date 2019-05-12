// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-12 16:56:18 B42348                 cmdx/[sort_file_lines_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in sort_file_lines.go use:
//      go test --run Test_sfln_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_sfln_sortFileLines_
func Test_sfln_sortFileLines_(t *testing.T) {
	zr.TBegin(t)
	// sortFileLines(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		sortFileLines(cmd, args)
	}
	test(Command{}, []string{})
} //                                                    Test_sfln_sortFileLines_

//end
