// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-28 14:06:54 094085                        [cmdx/replace_lines.go]
// -----------------------------------------------------------------------------

package main

import "sort"        // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zircon-Go

//  replaceLines(
//      lines []string,
//      finds [][]string,
//      repls [][]string,
//      caseMode zr.CaseMode,
//  ) (
//      changedLines []string,
//      changes int,
//  )
//
// # Subfunctions
//   (M replaceLinesM) getBatches(
//                         finds, repls []Lines,
//                     ) (
//                         []int, map[int]*FindReplLinesBatch,
//                     )
//   (M replaceLinesM) getTree(
//                         finds, repls []Lines,
//                         caseMode zr.CaseMode,
//                     ) (
//                         ret FindReplLinesTree,
//                     )
//   (M replaceLinesM) replaceMany(
//                         lines []string,
//                         finds []Lines,
//                         repls []Lines,
//                         caseMode zr.CaseMode,
//                     ) (
//                         changedLines Lines,
//                         changes int,
//                     )

// replaceLinesM joins all subfunctions used by replaceLines(),
// so that their names don't clutter the project's namespace.
type replaceLinesM struct{}

// replaceLines replaces multiple blocks of lines.
func replaceLines(
	lines []string,
	finds [][]string,
	repls [][]string,
	caseMode zr.CaseMode,
) (
	changedLines []string,
	changes int,
) {
	if len(finds) == 0 {
		return lines, 0 // avoid allocation
	}
	if DebugReplaceLines {
		zr.DV("--------------------------------------------------")
		zr.DV("replaceLines() args:")
		zr.DV("lines", lines)
		zr.DV("finds", finds)
		zr.DV("repls", repls)
		zr.DV("caseMode", caseMode)
		zr.DV("--------------------------------------------------")
	}
	// validate arguments
	if len(finds) != len(repls) {
		zr.Error(zr.EInvalidArg, ": lineCounts don't match:",
			len(finds), "and", len(repls))
		return Lines{}, 0 // erv
	}
	if caseMode != zr.IgnoreCase && caseMode != zr.MatchCase {
		zr.Error(zr.EInvalidArg,
			"^caseMode", ":", caseMode, "defaulting to 'MatchCase'")
		caseMode = zr.MatchCase
	}
	// copy [][]string to []Lines. Lis there a way to cast?
	var findLines = make([]Lines, len(finds))
	var replLines = make([]Lines, len(repls))
	for i, find := range finds {
		findLines = append(findLines, find)
		replLines = append(replLines, repls[i])
	}
	// make replacements by batches with largest number of lines first
	var M replaceLinesM
	var descLineCounts, batches = M.getBatches(findLines, replLines)
	for _, batchSize := range descLineCounts {
		var b, _ = batches[batchSize]
		var n = 0
		lines, n = M.replaceMany(lines, b.FindLines, b.ReplLines, caseMode)
		changes += n
	}
	if DebugReplaceLines {
		zr.DV("--------------------------------------------------")
		zr.DV("replaceLines() returns:")
		zr.DV("lines", lines)
		zr.DV("changes", changes)
		zr.DV("--------------------------------------------------")
	}
	return lines, changes
} //                                                                replaceLines

// -----------------------------------------------------------------------------
// # Subfunctions

// getBatches __
func (M replaceLinesM) getBatches(
	finds, repls []Lines,
) (
	[]int, map[int]*FindReplLinesBatch,
) {
	var lineCounts []int
	var batches = map[int]*FindReplLinesBatch{}
	for i, find := range finds {
		var n = len(find)
		var b, has = batches[n]
		if !has {
			batches[n] = &FindReplLinesBatch{}
			lineCounts = append(lineCounts, n)
		}
		b, _ = batches[n]
		b.FindLines = append(b.FindLines, find)
		b.ReplLines = append(b.ReplLines, repls[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lineCounts)))
	return lineCounts, batches
} //                                                                  getBatches

// getTree builds a tree from a slice of strings to find and to replace.
// The branches of the tree are made from each line in batches of repls.
func (M replaceLinesM) getTree(
	finds, repls []Lines,
	caseMode zr.CaseMode,
) (
	ret FindReplLinesTree,
) {
	ret.Sub = make(map[string]*FindReplLinesTree)
	for f, find := range finds {
		var node = &ret
		var last = len(find) - 1
		for i, line := range find {
			line = str.Trim(line, SPACES)
			if caseMode == zr.IgnoreCase {
				line = str.ToLower(line)
			}
			var _, exist = node.Sub[line]
			var sub *FindReplLinesTree
			if exist {
				sub = node.Sub[line]
			} else {
				sub = &FindReplLinesTree{
					Sub: make(map[string]*FindReplLinesTree),
				}
			}
			if i == last {
				sub.FindLines = find
				sub.ReplLines = repls[f]
			}
			node.Sub[line] = sub
			node = sub
		}
	}
	return ret
} //                                                                     getTree

// replaceMany __
func (M replaceLinesM) replaceMany(
	lines []string,
	finds []Lines,
	repls []Lines,
	caseMode zr.CaseMode,
) (
	changedLines Lines,
	changes int,
) {
	var linesLen = len(lines)
	var root = M.getTree(finds, repls, caseMode)
	var node = &root // *tree pointing to current branch
	var match = 0    // <- number of matching characters
	var prev = 0
	var ret []string
	for i := 0; i < linesLen; i++ {
		var line = str.Trim(lines[i], SPACES)
		if caseMode == zr.IgnoreCase {
			line = str.ToLower(line)
		}
		// check if the tree's branch has a key matching the current line
		// if not, reset matching count and start over from root
		{
			var sub, found = node.Sub[line]
			if !found {
				node = &root
				i -= match
				match = 0
				continue
			}
			match++
			node = sub
		}
		// even if found, keep matching if there are more matches to make.
		// when node.FindLines is not empty, it's a leaf node, then proceed
		var findLen = len(node.FindLines)
		if findLen == 0 || findLen != match {
			continue
		}
		// replacement starts:
		changes++
		// concatenate preceding lines that remain unchanged
		if prev < i-findLen+1 {
			ret = append(ret, lines[prev:i-findLen+1]...)
		}
		// append the replacement lines
		ret = append(ret, node.ReplLines...)
		prev = i + 1
		node = &root // restart matching from beginning
		match = 0
	}
	// write any remaining unchanged lines
	if prev < linesLen {
		ret = append(ret, lines[prev:]...)
	}
	return ret, changes
} //                                                                 replaceMany

//end
