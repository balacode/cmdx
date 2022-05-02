// -----------------------------------------------------------------------------
// CMDX Utilities Suite                             cmdx/[replace_lines_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

//  to test all items in replace_lines.go use:
//      go test --run Test_rpln_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"testing"

	"github.com/balacode/zr"
)

// go test --run Test_rpln_replaceLines_
func Test_rpln_replaceLines_(t *testing.T) {
	zr.TBegin(t)
	// replaceLines(
	//     lines []string,
	//     finds [][]string,
	//     repls [][]string,
	//     caseMode zr.CaseMode,
	// ) (changedLines []string, changes int)
	//
	test := func(
		// in:
		lines []string,
		finds [][]string,
		repls [][]string,
		caseMode zr.CaseMode,
		// out expected:
		changedLines []string,
		changes int, //
	) {
		changedLinesT, changesT := replaceLines(
			lines, finds, repls, caseMode,
		)
		zr.TEqual(t, changedLinesT, (changedLines))
		zr.TEqual(t, changesT, (changes))
	}
	test(
		[]string{}, [][]string{}, [][]string{}, zr.MatchCase,
		// out:
		[]string{}, 0,
	)
}

// end
