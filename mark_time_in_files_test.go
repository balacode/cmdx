// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 01:01:21 45B300              [cmdx/mark_time_in_files_test.go]
// -----------------------------------------------------------------------------

package main

// # Command Handler
//   Test_mtif_markTimeInFiles_(t *testing.T)
//
// # Support (File Scope)
//   Test_mtif_checksum_(t *testing.T)
//   Test_mtif_getTimeLogPath(t *testing.T)
//   Test_mtif_processDir_(t *testing.T)
//   Test_mtif_processFile_(t *testing.T)
//   Test_mtif_replaceVersion_(t *testing.T)

/*
to test all items in mark_time_in_files.go use:
    go test --run Test_mtif_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "time"    // standard
import "testing" // standard

import "github.com/balacode/zr" // Zirconium

// -----------------------------------------------------------------------------
// # Command Handler

// go test --run Test_mtif_markTimeInFiles_
func Test_mtif_markTimeInFiles_(t *testing.T) {
	zr.TBegin(t)
	// markTimeInFiles(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		markTimeInFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                                  Test_mtif_markTimeInFiles_

// -----------------------------------------------------------------------------
// # Support (File Scope)

// go test --run Test_mtif_checksum_
func Test_mtif_checksum_(t *testing.T) {
	zr.TBegin(t)
	// checksum(s string) string
	//
	var test = func(
		// in:
		s string,
		// out:
		ret string,
	) {
		var retT = checksum(s)
		zr.TEqual(t, retT, (ret))
	}
	test("",
		// out:
		"")
} //                                                         Test_mtif_checksum_

// go test --run Test_mtif_getTimeLogPath
func Test_mtif_getTimeLogPath(t *testing.T) {
	zr.TBegin(t)
	// getTimeLogPath(path)
	//
	// these paths and filenames don't need to physically exist
	// the function only processes strings, which are all specified here
	var oldPaths = TimeLogPaths
	TimeLogPaths = []string{
		`x:\tests`,
		`x:\tests\sub`,
		`x:\tests\sub\p1`,
		`x:\tests\sub\p2`,
		`x:\tests\sub\p3`,
		`x:\tests\other`,
	}
	var fn = getTimeLogPath
	zr.TEqual(t, fn(`x:\tests\sub\p2\main.go`),
		`x:\tests\sub\p2`,
	)
	zr.TEqual(t, fn(`x:\tests\file.txt`),
		`x:\tests`,
	)
	TimeLogPaths = oldPaths
} //                                                    Test_mtif_getTimeLogPath

// go test --run Test_mtif_processDir_
func Test_mtif_processDir_(t *testing.T) {
	zr.TBegin(t)
	// processDir(dir string, changeTime bool)
	//
	var test = func(
		// in:
		dir string, changeTime bool,
	) {
		processDir(dir, changeTime)
	}
	test("", false)
} //                                                       Test_mtif_processDir_

// go test --run Test_mtif_processFile_
func Test_mtif_processFile_(t *testing.T) {
	zr.TBegin(t)
	// processFile(path, name string, modTime time.Time) error
	//
	var test = func(
		// in:
		path, name string, modTime time.Time,
		// out expected:
		err error,
	) {
		var errT = processFile(path, name, modTime)
		zr.TEqual(t, errT, (err))
	}
	test("", "", time.Time{},
		// out:
		nil)
} //                                                      Test_mtif_processFile_

// go test --run Test_mtif_replaceVersion_
func Test_mtif_replaceVersion_(t *testing.T) {
	zr.TBegin(t)
	// replaceVersion(s, path, filename string, modTime time.Time) string
	//
	var test = func(
		// in:
		s, path, filename string, modTime time.Time,
		// out expected:
		ret string,
	) {
		var retT = replaceVersion(s, path, filename, modTime)
		zr.TEqual(t, retT, (ret))
	}
	test("", "", "", time.Time{},
		// out:
		"")
} //                                                   Test_mtif_replaceVersion_

//end