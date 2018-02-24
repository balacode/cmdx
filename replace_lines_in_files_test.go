// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:47 EC0E0C          [cmdx/replace_lines_in_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in replace_lines_in_files.go use:
    go test --run Test_rlif_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zirconium

// go test --run Test_rlif_replaceLinesInFiles_
func Test_rlif_replaceLinesInFiles_(t *testing.T) {
	zr.TBegin(t)
	// replaceLinesInFiles(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		replaceLinesInFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                              Test_rlif_replaceLinesInFiles_

//end
