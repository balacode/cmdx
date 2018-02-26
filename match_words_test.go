// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 14:45:21 7F36C2                     [cmdx/match_words_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in match_words.go use:
    go test --run Test_mtcw_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zircon-Go

// go test --run Test_mtcw_matchWords_
func Test_mtcw_matchWords_(t *testing.T) {
	zr.TBegin(t)
	// matchWords(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		matchWords(cmd, args)
	}
	test(Command{}, []string{})
} //                                                       Test_mtcw_matchWords_

//end
