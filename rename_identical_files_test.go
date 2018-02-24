// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:47 7D2553          [cmdx/rename_identical_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in rename_identical_files.go use:
    go test --run Test_ridf_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zirconium

// go test --run Test_ridf_renameIdenticalFiles_
func Test_ridf_renameIdenticalFiles_(t *testing.T) {
	zr.TBegin(t)
	// renameIdenticalFiles(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		renameIdenticalFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                             Test_ridf_renameIdenticalFiles_

//end
