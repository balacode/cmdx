// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-03-18 01:07:59 81C9E2               cmdx/[delete_identical_files.go]
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
		toFiles := toFilesMap[size]
		if len(toFiles) == 0 {
			continue
		}
		for _, from := range fromFiles {
			fromData, done := env.ReadFile(from.Path)
			if !done {
				continue
			}
			for i, to := range toFiles {
				if to.Size < 0 {
					continue
				}
				toData, done := env.ReadFile(to.Path)
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
