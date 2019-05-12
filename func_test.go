// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-12 16:56:18 5B1B38                            cmdx/[func_test.go]
// -----------------------------------------------------------------------------

package main

//  to test all items in func.go use:
//      go test --run Test_func_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_func_getFilesMap_
func Test_func_getFilesMap_(t *testing.T) {
	zr.TBegin(t)
	// getFilesMap(dir, filter string) FilesMap
	//
	test := func(
		// in:
		dir, filter string,
		// out expected:
		ret FilesMap) {
		retT := getFilesMap(dir, filter)
		zr.TEqual(t, retT, (ret))
	}
	test("", "",
		// out:
		FilesMap{})
} //                                                      Test_func_getFilesMap_

// go test --run Test_func_splitArgsFilter_
func Test_func_splitArgsFilter_(t *testing.T) {
	zr.TBegin(t)
	// splitArgsFilter(args []string) (retArgs []string, filter string)
	//
	test := func(
		// in:
		args []string,
		// out expected:
		retArgs []string, filter string,
	) {
		retArgsT, filterT := splitArgsFilter(args)
		zr.TEqual(t, retArgsT, (retArgs))
		zr.TEqual(t, filterT, (filter))
	}
	test([]string{},
		// out:
		[]string{}, "")
} //                                                  Test_func_splitArgsFilter_

// go test --run Test_func_trim_
func Test_func_trim_(t *testing.T) {
	zr.TBegin(t)
	// trim(s string) string
	//
	test := func(s string, ret string) {
		retT := trim(s)
		zr.TEqual(t, retT, (ret))
	}
	test("",
		// out:
		"")
} //                                                             Test_func_trim_

//end
