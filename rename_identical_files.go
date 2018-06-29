// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-09 01:03:17 E9F631               cmdx/[rename_identical_files.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"path/filepath"
)

// renameIdenticalFiles __
func renameIdenticalFiles(cmd Command, args []string) {
	if len(args) == 2 {
		env.Println(`
--------------------------------------------------------------------------------
'rdup' or 'ren-dup' command: Rename identical files
--------------------------------------------------------------------------------

Carries out very rapid bulk renaming of files.
Requires a <source> and <target> folder.

Reads the names and sizes of all files in <source> and its subfolders,
then does the same for <target>, then compares the content of each
file with the same size. If any file in <target> has the same content
as a file in <source>, renames it to the file's name in <source>.

What is this used for? You can use this to organize media files, etc,
for example, if you have some photos that you named in one
folder and want to rename all matching files in another folder.

Note: this command doesn't filter file extensions and affects all matching
files in <target> (the <source> is never changed). Run it with care.
--------------------------------------------------------------------------------
`)
		return
	}
	var filter string
	args, filter = splitArgsFilter(args)
	if len(args) < 2 {
		env.Println("'rename' requires: <source dir> and <target dir>")
		return
	}
	var toFilesMap = getFilesMap(args[1], filter)
	for size, fromFiles := range getFilesMap(args[0], filter) {
		var toFiles = toFilesMap[size]
		if len(toFiles) == 0 {
			continue
		}
		for _, from := range fromFiles {
			var fromName = filepath.Base(from.Path)
			var fromData, done = env.ReadFile(from.Path)
			if !done {
				continue
			}
			for i, to := range toFiles {
				var toName = filepath.Base(to.Path)
				if toName == fromName {
					continue
				}
				var toData, done = env.ReadFile(to.Path)
				if !done {
					continue
				}
				if !bytes.Equal(fromData, toData) {
					continue
				}
				var renamedPath = filepath.Dir(to.Path) +
					env.PathSeparator() + fromName
				env.RenameFile(to.Path, renamedPath)
				env.Println("renamed", to.Path, " -> ", renamedPath)
				toFiles[i].Path = renamedPath
			}
		}
	}
} //                                                        renameIdenticalFiles

//end
