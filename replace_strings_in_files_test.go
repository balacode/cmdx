// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 14:45:21 F63211        [cmdx/replace_strings_in_files_test.go]
// -----------------------------------------------------------------------------

package main

/*
to test all items in replace_strings_in_files.go use:
    go test --run Test_rsif_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "sync"    // standard
import "testing" // standard

import "github.com/balacode/zr" // Zircon-Go

// go test --run Test_rsif_replaceStringsInFiles_
func Test_rsif_replaceStringsInFiles_(t *testing.T) {
	zr.TBegin(t)
	// replaceStringsInFiles(cmd Command, args []string)
	//
	var test = func(
		// in:
		cmd Command, args []string,
	) {
		replaceStringsInFiles(cmd, args)
	}
	test(Command{}, []string{})
} //                                            Test_rsif_replaceStringsInFiles_

// go test --run Test_rsif_replaceAsync_
func Test_rsif_replaceAsync_(t *testing.T) {
	zr.TBegin(t)
	// replaceAsync(task *sync.WaitGroup, configFile string, cmd ReplCmd)
	//
	var test = func(
		// in:
		task *sync.WaitGroup, configFile string, cmd ReplCmd,
	) {
		replaceAsync(task, configFile, cmd)
	}
	var cmd = ReplCmd{}
	test(nil, "", cmd)
} //                                                     Test_rsif_replaceAsync_

// go test --run Test_rsif_replaceFileAsync_
func Test_rsif_replaceFileAsync_(t *testing.T) {
	//  replaceFileAsync(
	//  	task *sync.WaitGroup,
	//  	configFile string,
	//  	filename string,
	//  	content string,
	//  	items []ReplItem,
	//  )
	zr.TBegin(t)
	var test = func(
		// in:
		task *sync.WaitGroup,
		configFile string,
		filename string,
		content string,
		items []ReplItem,
	) {
		replaceFileAsync(task, configFile, filename, content, items)
	}
	test(nil, "", "", "", []ReplItem{})
} //                                                 Test_rsif_replaceFileAsync_

//end
