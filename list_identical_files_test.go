// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 14:45:21 E350B9            [cmdx/list_identical_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in list_identical_files.go use:
    go test --run Test_lsif_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zircon-Go

// go test --run Test_lsif_listIdenticalFiles_
func Test_lsif_listIdenticalFiles_(t *testing.T) {
	zr.TBegin(t)
	// listIdenticalFiles(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		listIdenticalFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                               Test_lsif_listIdenticalFiles_

//end
