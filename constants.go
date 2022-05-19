// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                      cmdx/[constants.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// FindInFilesDialog _ _
const FindInFilesDialog = `
MRK  ---------------------------------------------------------------------------
MRK |  FIND IN FILES:    (to find multiple strings place them on multiple lines)
MRK  ---------------------------------------------------------------------------

MRK  ---------------------------------------------------------------------------
MRK |  find in path: X:/path/
MRK |  subfolders:   y
MRK |
MRK |  ignore case:  y      mark in source (prefix):
MRK |  whole word:   y      mark in source (suffix):
MRK |
MRK |  show in result:
MRK |  path          y      preceding lines: 0
MRK |  filename      y      following lines: 0
MRK |  line number   y
MRK |  line text     y      paste results here: n
MRK |                                                                 WAITING...
MRK  ---------------------------------------------------------------------------
`

// Spaces is a string of all white-space characters,
// which includes spaces, tabs, and newline characters.
const Spaces = " \a\b\f\n\r\t\v"

// ExtraTextFileExts contains additional text file
// extensions not listed in (zr-fs) fs.TextFileExts
var ExtraTextFileExts = []string{
	//
	// Go module files
	"mod",
	"sum",
	//
	// YAML files // TODO: remove after upgrading zr/fs
	"yml",
	"yaml",
	//
	// Google Protocol Buffer files
	"proto",
}

// end
