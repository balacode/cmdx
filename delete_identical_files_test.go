// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-06 06:03:40 1BE940          cmdx/[delete_identical_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in delete_identical_files.go use:
    go test --run Test_dlif_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_dlif_deleteIdenticalFiles_
func Test_dlif_deleteIdenticalFiles_(t *testing.T) {
	zr.TBegin(t)
	// deleteIdenticalFiles(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		deleteIdenticalFiles(cmd, args)
		// TODO: implement Test_dlif_deleteIdenticalFiles_()
	}
	test(Command{}, []string{})
} //                                             Test_dlif_deleteIdenticalFiles_

//end
