// -----------------------------------------------------------------------------
// CMDX Utility                             cmdx/[mark_errors_in_source_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

//  to test all items in mark_errors_in_source.go use:
//      go test --run Test_meis_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

//  to test this manually:
//      remove any errors, and compile cx.exe:
//          build.bat
//      add some errors in source, then:
//          go build -gcflags="-e" -o tmp.exe 2> build.log
//          cx mark-errors -buildlog=.\build.log
//          del tmp.exe

// -----------------------------------------------------------------------------
// # Command Handler

// go test --run Test_meis_markErrorsInSource_
func Test_meis_markErrorsInSource_(t *testing.T) {
	zr.TBegin(t)
	// markErrorsInSource(cmd Command, args []string)
	//
	test := func(
		// in:
		cmd Command, args []string,
	) {
		markErrorsInSource(cmd, args)
		// TODO: implement unit test
	}
	test(Command{}, []string{})
} //                                               Test_meis_markErrorsInSource_

// -----------------------------------------------------------------------------
// # Support (File Scope)

// go test --run Test_meis_isErrorComment_
func Test_meis_isErrorComment_(t *testing.T) {
	zr.TBegin(t)
	// isErrorComment(line string) bool
	//
	test := func(
		// in:
		line string,
		// out expected:
		ret bool,
	) {
		retT := isErrorComment(line)
		zr.TEqual(t, retT, (ret))
	}
	test("",
		// out:
		false)
} //                                                   Test_meis_isErrorComment_

// go test --run Test_meis_makePath_
func Test_meis_makePath_(t *testing.T) {
	zr.TBegin(t)
	// makePath(absPath, relPath string) string
	//
	test := func(
		// in:
		absPath, relPath string,
		// out expected:
		ret string,
	) {
		retT := makePath(absPath, relPath)
		zr.TEqual(t, retT, (ret))
	}
	test("", "",
		// out:
		"")
} //                                                         Test_meis_makePath_

// go test --run Test_meis_readBuildIssues_
func Test_meis_readBuildIssues_(t *testing.T) {
	zr.TBegin(t)
	// readBuildIssues(buildLog string) (ret []BuildIssue)
	//
	test := func(
		// in:
		buildLog string,
		// out expected:
		ret []BuildIssue,
	) {
		retT := readBuildIssues(buildLog)
		zr.TEqual(t, retT, (ret))
	}
	test("",
		// out:
		[]BuildIssue{})
} //                                                  Test_meis_readBuildIssues_

// go test --run Test_meis_removeOldErrorComments_
func Test_meis_removeOldErrorComments_(t *testing.T) {
	zr.TBegin(t)
	// removeOldErrorComments(lines []string) []string
	//
	test := func(
		// in:
		lines []string,
		// out expected:
		ret []string,
	) {
		retT := removeOldErrorComments(lines)
		zr.TEqual(t, retT, (ret))
	}
	test([]string{},
		// out:
		[]string{})
} //                                           Test_meis_removeOldErrorComments_

// go test --run Test_meis_saveFile_
func Test_meis_saveFile_(t *testing.T) {
	zr.TBegin(t)
	// saveFile(buildPath, filename string, lines []string)
	//
	test := func(
		// in:
		buildPath, filename string, lines []string,
	) {
		saveFile(buildPath, filename, lines)
	}
	test("", "", []string{})
} //                                                         Test_meis_saveFile_

// end
