// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:38:12 0DAF5E                 [cmdx/rename_files_by_hash.go]
// -----------------------------------------------------------------------------

package main

import "path/filepath" // standard
import str "strings"   // standard

import "github.com/balacode/zr" // Zirconium
import "ase/zr/whirl"           // Zirconium

// renameFilesByHash __
func renameFilesByHash(cmd Command, args []string) {
	var filter string
	args, filter = splitArgsFilter(args)
	if len(args) == 0 {
		env.Println(`
--------------------------------------------------------------------------------
'rh' or 'ren-hash' command: Rename files to hash
--------------------------------------------------------------------------------
--------------------------------------------------------------------------------
`)
		return
	}
	if len(args) < 1 {
		env.Println("'rename' requires: <source dir>")
		return
	}
	for _, files := range getFilesMap(args[0], filter) {
		for i, file := range files {
			_ = i // <- unused size
			var name = filepath.Base(file.Path)
			var data, done = env.ReadFile(file.Path)
			if !done {
				continue
			}
			var hash = zr.HexStringOfBytes(zr.FoldXorBytes(
				whirl.HashOfBytes(data, []byte{}), 4,
			))
			hash = str.ToLower(hash)
			// skip filenames that already contain the hash
			if str.Contains(str.ToLower(name), hash) {
				continue
			}
			var newName = hash + "." + name
			var newPath = filepath.Dir(file.Path) +
				env.PathSeparator() +
				newName
			env.RenameFile(file.Path, newPath)
			env.Println("renamed", file.Path, " -> ", newName)
			files[i].Path = newPath
		}
	}
} //                                                           renameFilesByHash

//end
