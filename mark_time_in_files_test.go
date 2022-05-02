// -----------------------------------------------------------------------------
// CMDX Utilities Suite                        cmdx/[mark_time_in_files_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// # Command Handler
//   Test_mtif_markTimeInFiles_(t *testing.T)
//
// # Support (File Scope)
//   Test_mtif_getTimeLogPath(t *testing.T)
//   Test_mtif_processDir_(t *testing.T)
//   Test_mtif_processFile_(t *testing.T)
//   Test_mtif_replaceVersion_(t *testing.T)

//  to test all items in mark_time_in_files.go use:
//      go test --run Test_mtif_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"
	"time"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Command Handler

// go test --run Test_mtif_markTimeInFiles_
func Test_mtif_markTimeInFiles_(t *testing.T) {
	zr.TBegin(t)
	// markTimeInFiles(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		markTimeInFiles(cmd, args)
	}
	test(Command{}, []string{})
}

// -----------------------------------------------------------------------------
// # Support (File Scope)

// go test --run Test_mtif_getTimeLogPath
func Test_mtif_getTimeLogPath(t *testing.T) {
	zr.TBegin(t)
	// getTimeLogPath(path)
	//
	// these paths and filenames don't need to physically exist
	// the function only processes strings, which are all specified here
	oldPaths := TimeLogPaths
	TimeLogPaths = []string{
		`X:\tests`,
		`X:\tests\sub`,
		`X:\tests\sub\p1`,
		`X:\tests\sub\p2`,
		`X:\tests\sub\p3`,
		`X:\tests\other`,
	}
	fn := getTimeLogPath
	zr.TEqual(t, fn(`X:\tests\sub\p2\main.go`),
		`X:\tests\sub\p2`,
	)
	zr.TEqual(t, fn(`X:\tests\file.txt`),
		`X:\tests`,
	)
	TimeLogPaths = oldPaths
}

// go test --run Test_mtif_processDir_
func Test_mtif_processDir_(t *testing.T) {
	zr.TBegin(t)
	// processDir(dir string, changeTime bool)
	//
	test := func(
		// in:
		dir string, changeTime bool,
	) {
		processDir(dir, changeTime)
	}
	test("", false)
}

// go test --run Test_mtif_processFile_
func Test_mtif_processFile_(t *testing.T) {
	zr.TBegin(t)
	// processFile(path, name string, modTime time.Time) error
	//
	test := func(
		// in:
		path, name string, modTime time.Time,
		// out expected:
		err error,
	) {
		errT := processFile(path, name, modTime)
		zr.TEqual(t, errT, (err))
	}
	test("", "", time.Time{},
		// out:
		nil)
}

// go test --run Test_mtif_replaceVersion_
func Test_mtif_replaceVersion_(t *testing.T) {
	zr.TBegin(t)
	// replaceVersion(s, path, filename string, modTime time.Time) string
	//
	test := func(
		// in:
		s, path, filename string, modTime time.Time,
		// out expected:
		ret string,
	) {
		retT := replaceVersion(s, path, filename, modTime)
		zr.TEqual(t, retT, (ret))
	}
	test("", "", "", time.Time{},
		// out:
		"")
}

// end
