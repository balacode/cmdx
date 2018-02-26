// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-02-26 14:45:21 85B28C                         [cmdx/replace_time.go]
// -----------------------------------------------------------------------------

package main

import "bytes"       // standard
import "regexp"      // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zircon-Go

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
			var content = str.Trim(string(data), zr.SPACES)
			ar = str.Split(content, zr.LF)
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
		var s = str.Trim(string(data), zr.SPACES)
		toLines = str.Split(s, zr.LF)
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
						out.WriteString(s + zr.LF)
					}
					continue
				}
			}
			out.WriteString(s + zr.LF)
		}
	}
	if !env.WriteFile(toFile, out.Bytes()) {
		return
	}
	env.Printf("written '%s'"+zr.LF, toFile)
} //                                                                 replaceTime

//end
