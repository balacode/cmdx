// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2021-02-10 09:29:15 42EF3D          cmdx/[replace_lines_in_files_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in replace_lines_in_files.go use:
//      go test --run Test_rlif_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_rlif_replaceLinesInFiles_
func Test_rlif_replaceLinesInFiles_(t *testing.T) {
	zr.TBegin(t)
	// replaceLinesInFiles(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		replaceLinesInFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                              Test_rlif_replaceLinesInFiles_

// end
