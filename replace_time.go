// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-11 04:25:01 8C1362                         cmdx/[replace_time.go]
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"regexp"
	"strings"
)

// replaceTime __
func replaceTime(cmd Command, args []string) {
	if len(args) < 2 {
		env.Println("'replace-time' requires: <source file> and <target file>")
		return
	}
	var (
		fromFile  = args[0]
		toFile    = args[1]
		fromLines = map[string][]string{}
		toLines   []string
		validTime = regexp.MustCompile( // YYYY-MM-DD hh:mm
			"^[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2} ")
	)
	{ // fill 'fromLines' array
		var ar []string
		{
			data, done := env.ReadFile(fromFile)
			if !done {
				return
			}
			content := strings.TrimSpace(string(data))
			ar = strings.Split(content, "\n")
		}
		for _, s := range ar {
			if len(s) < 16 {
				continue
			}
			if !validTime.MatchString(s) {
				continue
			}
			tm := s[:16]
			fromLines[tm] = append(fromLines[tm], s)
		}
	}
	{ // fill 'toLines' array
		data, done := env.ReadFile(toFile)
		if !done {
			return
		}
		s := strings.TrimSpace(string(data))
		toLines = strings.Split(s, "\n")
	}
	var out bytes.Buffer
	{
		var tmPrev string
		for _, s := range toLines {
			if validTime.MatchString(s) {
				tm := s[:16]
				if tm == tmPrev {
					continue
				}
				tmPrev = tm
				if from, exist := fromLines[tm]; exist {
					for _, s := range from {
						out.WriteString(s + "\n")
					}
					continue
				}
			}
			out.WriteString(s + "\n")
		}
	}
	if !env.WriteFile(toFile, out.Bytes()) {
		return
	}
	env.Printf("written '%s'\n", toFile)
} //                                                                 replaceTime

//end
