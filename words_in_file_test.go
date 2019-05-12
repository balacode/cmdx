// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-12 16:56:18 D982D5                   cmdx/[words_in_file_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in words_in_file.go use:
//      go test --run Test_wdif_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_wdif_wordsInFile_
func Test_wdif_wordsInFile_(t *testing.T) {
	zr.TBegin(t)
	// wordsInFile(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		wordsInFile(cmd, args)
	}
	test(Command{}, []string{})
} //                                                      Test_wdif_wordsInFile_

//end
