// -----------------------------------------------------------------------------
// CMDX Utilities Suite                      cmdx/[list_identical_files_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

//  to test all items in list_identical_files.go use:
//      go test --run Test_lsif_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_lsif_listIdenticalFiles_
func Test_lsif_listIdenticalFiles_(t *testing.T) {
	zr.TBegin(t)
	// listIdenticalFiles(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		listIdenticalFiles(cmd, args)
	}
	test(Command{}, []string{})
}

// end
