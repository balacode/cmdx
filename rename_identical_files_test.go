// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-03-18 01:07:59 56C617          cmdx/[rename_identical_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in rename_identical_files.go use:
    go test --run Test_ridf_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_ridf_renameIdenticalFiles_
func Test_ridf_renameIdenticalFiles_(t *testing.T) {
	zr.TBegin(t)
	// renameIdenticalFiles(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		renameIdenticalFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                             Test_ridf_renameIdenticalFiles_

//end
