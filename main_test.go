// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:46 7235D2                            [cmdx/main_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in main.go use:
    go test --run Test_main_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "testing" // standard

import "github.com/balacode/zr" // Zirconium

// -----------------------------------------------------------------------------
// # Main Function

// go test --run Test_main_
func Test_main_(t *testing.T) {
	zr.TBegin(t)
	// main()
	//
	var test = func() {
		//TODO: main()
	}
	test()
} //                                                                  Test_main_

//end
