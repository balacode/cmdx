// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-28 14:22:33 F92B1A                 [cmdx/rename_files_by_hash.go]
// -----------------------------------------------------------------------------

package main

import "crypto/sha512" // standard
import "path/filepath" // standard
import str "strings"   // standard

import "github.com/balacode/zr" // Zircon-Go

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
				hashOfBytes(data, []byte{}), 4,
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

// hashOfBytes returns the SHA-512 hash of a byte slice.
// It also requires a 'salt' argument.
func hashOfBytes(ar []byte, salt []byte) []byte {
	var input []byte
	input = append(input, salt[:]...)
	input = append(input, ar...)
	var hash = sha512.Sum512(input)
	return hash[:]
} //                                                                 hashOfBytes

//end
