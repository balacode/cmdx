// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-24 00:37:47 DC6B0F                                 [cmdx/type.go]
// -----------------------------------------------------------------------------

package main

import "time" // standard

import "github.com/balacode/zr" // Zirconium

// -----------------------------------------------------------------------------
// # Types

// BuildIssue stores the message and location of each build issue/error
type BuildIssue struct {
	File string // file name (may be with or without a path)
	Line int    // line number
	Col  int    // column number
	Msg  string // message (usually a build error message)
} //                                                                  BuildIssue

// FilesMap groups files by their file sizes.
// (If two files have different sizes, they're definitely
// different, so there's no need to open and compare them.)
type FilesMap map[int64][]*PathAndSize

// FindReplLinesBatch holds groups of searches
// and corresponding replacements.
type FindReplLinesBatch struct {
	FindLines []Lines
	ReplLines []Lines
} //                                                          FindReplLinesBatch

// FindReplLines __
type FindReplLines struct {
	Path      string
	Exts      []string
	FindLines []string
	ReplLines []string
	CaseMode  zr.CaseMode
} //                                                               FindReplLines

// FindReplLinesTree __
type FindReplLinesTree struct {
	FindLines Lines                         // what to find
	ReplLines Lines                         // what to replace with
	Sub       map[string]*FindReplLinesTree // a 'branch' of the tree
} //                                                           FindReplLinesTree

// Lines __
type Lines []string

// PathAndSize stores a fully-qualified file name and the file's size.
type PathAndSize struct {
	Path string
	Size int64
} //                                                                 PathAndSize

// ReplCmd __
type ReplCmd struct {
	Path  string
	Exts  []string
	Mark  string
	Items []ReplItem
} //                                                                     ReplCmd

// ReplItem __
type ReplItem struct {
	Find     string
	Repl     string
	CaseMode zr.CaseMode
	WordMode zr.WordMode
} //                                                                    ReplItem

// TextFile holds file details and content of a text file.
type TextFile struct {
	Path    string
	ModTime time.Time
	Size    int64
	Lines   []string
} //                                                                    TextFile

//end
