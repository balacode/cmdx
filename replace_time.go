// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-28 13:59:12 B6807D                         cmdx/[replace_time.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"regexp"
	str "strings"
)

// replaceTime __
func replaceTime(cmd Command, args []string) {
	if len(args) < 2 {
		env.Println("'replace-time' requires: <source file> and <target file>")
		return
	}
	var fromFile = args[0]
	var toFile = args[1]
	var validTime = regexp.MustCompile( // YYYY-MM-DD hh:mm
		"^[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2} ")
	var fromLines = map[string][]string{}
	var toLines []string
	{ // fill 'fromLines' array
		var ar []string
		{
			var data, done = env.ReadFile(fromFile)
			if !done {
				return
			}
			var content = str.Trim(string(data), SPACES)
			ar = str.Split(content, LF)
		}
		for _, s := range ar {
			if len(s) < 16 {
				continue
			}
			if !validTime.MatchString(s) {
				continue
			}
			var tm = s[:16]
			fromLines[tm] = append(fromLines[tm], s)
		}
	}
	{ // fill 'toLines' array
		var data, done = env.ReadFile(toFile)
		if !done {
			return
		}
		var s = str.Trim(string(data), SPACES)
		toLines = str.Split(s, LF)
	}
	var out bytes.Buffer
	{
		var tmPrev string
		for _, s := range toLines {
			if validTime.MatchString(s) {
				var tm = s[:16]
				if tm == tmPrev {
					continue
				}
				tmPrev = tm
				if from, exist := fromLines[tm]; exist {
					for _, s := range from {
						out.WriteString(s + LF)
					}
					continue
				}
			}
			out.WriteString(s + LF)
		}
	}
	if !env.WriteFile(toFile, out.Bytes()) {
		return
	}
	env.Printf("written '%s'"+LF, toFile)
} //                                                                 replaceTime

//end
