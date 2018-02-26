// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 14:45:21 C844BC            [cmdx/rename_files_by_hash_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in rename_files_by_hash.go use:
    go test --run Test_rfbh_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zircon-Go

// go test --run Test_rfbh_renameFilesByHash_
func Test_rfbh_renameFilesByHash_(t *testing.T) {
	zr.TBegin(t)
	// renameFilesByHash(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		renameFilesByHash(cmd, args)
	}
	test(Command{}, []string{})
} //                                                Test_rfbh_renameFilesByHash_

//end
