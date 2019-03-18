// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-03-18 01:07:59 FFE9C8                  cmdx/[replace_shared_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in replace_shared.go use:
    go test --run Test_rsha_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_rsha_getConfigBool_
func Test_rsha_getConfigBool_(t *testing.T) {
	zr.TBegin(t)
	// getConfigBool(s, keyword string) (value, exists bool)
	//
	test := func(
		// in:
		s, keyword string,
		// out:
		value, exists bool,
	) {
		valueT, existsT := getConfigBool(s, keyword)
		zr.TEqual(t, valueT, (value))
		zr.TEqual(t, existsT, (exists))
	}
	test("", "",
		// out:
		false, false)
} //                                                    Test_rsha_getConfigBool_

// go test --run Test_rsha_hasConfigBool_
func Test_rsha_hasConfigBool_(t *testing.T) {
	zr.TBegin(t)
	// hasConfigBool(s, keyword string) (ret bool)
	//
	test := func(s, keyword string, ret bool) {
		retT := hasConfigBool(s, keyword)
		zr.TEqual(t, retT, (ret))
	}
	test("", "",
		// out:
		false)
} //                                                    Test_rsha_hasConfigBool_

// go test --run Test_rsha_readConfigFileLines_
func Test_rsha_readConfigFileLines_(t *testing.T) {
	zr.TBegin(t)
	// readConfigFileLines(configFile string) (configLines []string)
	//
	test := func(
		// in:
		configFile string,
		// out:
		configLines []string,
	) {
		configLinesT := readConfigFileLines(configFile)
		zr.TEqual(t, configLinesT, (configLines))
	}
	test("",
		// out:
		[]string{})
} //                                              Test_rsha_readConfigFileLines_

//end
