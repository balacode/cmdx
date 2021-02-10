// -----------------------------------------------------------------------------
// CMDX Utility                               cmdx/[sort_lines_in_files_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
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

// end
