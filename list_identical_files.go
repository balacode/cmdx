// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-09 01:03:17 EBA18A                 cmdx/[list_identical_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
)

// listIdenticalFiles __
func listIdenticalFiles(cmd Command, args []string) {
	var filter string
	args, filter = splitArgsFilter(args)
	if len(args) < 1 {
		env.Println("'duplicates' requires:" +
			" <source dir>, or <source dir> and <target dir>")
		return
	}
	var toFilesMap FilesMap
	if len(args) == 1 {
		toFilesMap = getFilesMap(args[0], filter)
	} else {
		toFilesMap = getFilesMap(args[1], filter)
	}
	var duplicates = make(map[string]bool, 1)
	for size, fromFiles := range getFilesMap(args[0], filter) {
		var toFiles = toFilesMap[size]
		if len(toFiles) == 0 {
			continue
		}
		for _, from := range fromFiles {
			if duplicates[from.Path] {
				continue
			}
			var fromData []byte
			var first = true
			for _, to := range toFiles {
				if from.Path == to.Path {
					continue
				}
				if len(fromData) == 0 {
					var data, done = env.ReadFile(from.Path)
					if !done {
						continue
					}
					fromData = data
				}
				var toData, done = env.ReadFile(to.Path)
				if !done {
					continue
				}
				if !bytes.Equal(fromData, toData) {
					continue
				}
				if first {
					first = false
					env.Println()
					env.Println(from.Path)
				}
				env.Println(to.Path)
				duplicates[to.Path] = true
			}
		}
	}
} //                                                          listIdenticalFiles

//TODO: add a function to call when each duplicate is found.

//TODO: can be moved to 'fs' with some changes

//end
