// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-09 01:03:17 E97614               cmdx/[delete_identical_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
)

// deleteIdenticalFiles __
func deleteIdenticalFiles(cmd Command, args []string) {
	var filter string
	args, filter = splitArgsFilter(args)
	if len(args) < 2 {
		env.Println("'delete' requires: <source dir> and <target dir>")
		return
	}
	var toFilesMap map[int64][]*PathAndSize = getFilesMap(args[1], filter)
	for size, fromFiles := range getFilesMap(args[0], filter) {
		var toFiles = toFilesMap[size]
		if len(toFiles) == 0 {
			continue
		}
		for _, from := range fromFiles {
			var fromData, done = env.ReadFile(from.Path)
			if !done {
				continue
			}
			for i, to := range toFiles {
				if to.Size < 0 {
					continue
				}
				var toData, done = env.ReadFile(to.Path)
				if !done {
					continue
				}
				if !bytes.Equal(fromData, toData) {
					continue
				}
				if env.DeleteFile(to.Path) {
					toFiles[i].Size = -1
					env.Println("deleted", to.Path)
				}
			}
		}
	}
} //                                                        deleteIdenticalFiles

//end
