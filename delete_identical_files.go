// -----------------------------------------------------------------------------
// CMDX Utilities Suite                         cmdx/[delete_identical_files.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"bytes"

	"github.com/balacode/zr"
)

// deleteIdenticalFiles _ _
func deleteIdenticalFiles(cmd Command, args []string) {
	var filter string
	args, filter = splitArgsFilter(args)
	if len(args) < 2 {
		env.Println("'delete' requires: <source dir> and <target dir>")
		return
	}
	var (
		totalDeletedBytes = int64(0)
		totalDeletedFiles = 0
		toFilesMap        = getFilesMap(args[1], filter)
		//                  ^ map[int64][]*PathAndSize
	)
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
					totalDeletedBytes += size
					totalDeletedFiles++
				}
			}
		}
	}
	if totalDeletedBytes > 0 || totalDeletedFiles > 0 {
		nBytes := zr.ByteSizeString(totalDeletedBytes /*useSI=*/, false)
		s := zr.IfString(totalDeletedFiles == 1, "", "s")
		env.Println("deleted", nBytes, "in", totalDeletedFiles, "file"+s)
	}
}

// end
