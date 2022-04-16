// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                   cmdx/[is_text_file.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	fs "github.com/balacode/zr-fs"
)

// isTextFile returns true if the given file name
// represents a text file type. For example "readme.txt"
// returns true, while "image.png" returns false.
func isTextFile(filename string) bool {
	exts := append(env.TextFileExts(), ExtraTextFileExts...)
	ret := fs.IsFileExt(filename, exts)
	return ret
}

// end
