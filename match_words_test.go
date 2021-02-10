// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2021-02-10 09:29:15 94D41E                     cmdx/[match_words_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in match_words.go use:
//      go test --run Test_mtcw_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_mtcw_matchWords_
func Test_mtcw_matchWords_(t *testing.T) {
	zr.TBegin(t)
	// matchWords(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		matchWords(cmd, args)
	}
	test(Command{}, []string{})
} //                                                       Test_mtcw_matchWords_

// end
