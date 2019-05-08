// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-08 11:24:44 054CE4                         cmdx/[replace_time.go]
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
	fromFile := args[0]
	toFile := args[1]
	validTime := regexp.MustCompile( // YYYY-MM-DD hh:mm
		"^[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2} ")
	fromLines := map[string][]string{}
	var toLines []string
	{ // fill 'fromLines' array
		var ar []string
		{
			data, done := env.ReadFile(fromFile)
			if !done {
				return
			}
			content := strings.TrimSpace(string(data))
			ar = strings.Split(content, LF)
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
		toLines = strings.Split(s, LF)
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
