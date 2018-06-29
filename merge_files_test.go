// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-24 03:05:23 CD745B                     cmdx/[merge_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in merge_files.go use:
    go test --run Test_mgfl_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/zr"
)

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