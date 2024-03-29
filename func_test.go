// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                      cmdx/[func_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
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

// go test --run Test_func_checksum_
func Test_func_checksum_(t *testing.T) {
	zr.TBegin(t)
	// checksum(s string) string
	//
	test := func(
		// in:
		s string,
		// out:
		ret string,
	) {
		retT := checksum(s)
		zr.TEqual(t, retT, (ret))
	}
	test("",
		// out:
		"000000")
}

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
}

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
}

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
}

// end
