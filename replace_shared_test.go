// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-09 01:03:17 2BDEED                  [cmdx/replace_shared_test.go]
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

	"github.com/balacode/zr" // Zircon-Go
)

// go test --run Test_rsha_getConfigBool_
func Test_rsha_getConfigBool_(t *testing.T) {
	zr.TBegin(t)
	// getConfigBool(s, keyword string) (value, exists bool)
	//
	var test = func(
		// in:
		s, keyword string,
		// out:
		value, exists bool,
	) {
		var valueT, existsT = getConfigBool(s, keyword)
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
	var test = func(s, keyword string, ret bool) {
		var retT = hasConfigBool(s, keyword)
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
	var test = func(
		// in:
		configFile string,
		// out:
		configLines []string,
	) {
		var configLinesT = readConfigFileLines(configFile)
		zr.TEqual(t, configLinesT, (configLines))
	}
	test("",
		// out:
		[]string{})
} //                                              Test_rsha_readConfigFileLines_

//end
