// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-28 16:18:03 48C317                            cmdx/[constants.go]
// -----------------------------------------------------------------------------

package main

// FindInFilesDialog __
const FindInFilesDialog = `
MRK  ---------------------------------------------------------------------------
MRK |  FIND IN FILES:    (to find multiple strings place them on multiple lines)
MRK  ---------------------------------------------------------------------------

MRK  ---------------------------------------------------------------------------
MRK |  find in path: x:/path/
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

// CR is a string with a single carriage return
// character (decimal 13, hex 0D)
const CR = "\r"

// LF specifies a line feed string ("\n").
const LF = "\n"

// SPACES is a string of all white-space characters,
// which includes spaces, tabs, and newline characters.
const SPACES = " \a\b\f\n\r\t\v"

//eof
