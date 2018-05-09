// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-09 01:03:17 D8C173                    [cmdx/replace_time_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in replace_time.go use:
    go test --run Test_rptm_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/zr" // Zircon-Go
)

// go test --run Test_rptm_replaceTime_
func Test_rptm_replaceTime_(t *testing.T) {
	zr.TBegin(t)
	// replaceTime(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		replaceTime(cmd, args)
	}
	test(Command{}, []string{})
} //                                                      Test_rptm_replaceTime_

//end
