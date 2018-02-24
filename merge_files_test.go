// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:47 CBB070                     [cmdx/merge_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in merge_files.go use:
    go test --run Test_mgfl_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zirconium

// go test --run Test_mgfl_mergeFiles_
func Test_mgfl_mergeFiles_(t *testing.T) {
	zr.TBegin(t)
	// mergeFiles(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		mergeFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                                       Test_mgfl_mergeFiles_

//end
