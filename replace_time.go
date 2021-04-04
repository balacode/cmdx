// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                   cmdx/[replace_time.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"bytes"
	"regexp"
	"strings"
)

const TR_TIME_LEN = 16 // length of 'YYYY-MM-DD hh:mm' in characters

// replaceTime _ _
func replaceTime(cmd Command, args []string) {
	if len(args) < 2 {
		env.Println("'replace-time' requires: <source file> and <target file>")
		return
	}
	var (
		fromFile  = args[0]
		toFile    = args[1]
		validTime = regexp.MustCompile(`^\d{4}-\d\d-\d\d \d\d:\d\d`)
		//                             `YYYY-MM-DD hh:mm` ^
	)
	var fromLines = map[string][]string{}
	{
		var lines []string
		{
			data, done := env.ReadFile(fromFile)
			if !done {
				return
			}
			content := strings.TrimSpace(string(data))
			lines = strings.Split(content, "\n")
		}
		for _, line := range lines {
			if len(line) < TR_TIME_LEN {
				continue
			}
			if !validTime.MatchString(line) {
				continue
			}
			tm := line[:TR_TIME_LEN]
			fromLines[tm] = append(fromLines[tm], line)
		}
	}
	var toLines []string
	{
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
		for _, line := range toLines {
			if validTime.MatchString(line) {
				tm := line[:TR_TIME_LEN]
				from, exist := fromLines[tm]
				if exist {
					if tm == tmPrev {
						continue
					}
					tmPrev = tm
					for _, s := range from {
						out.WriteString(s + "\n")
					}
					continue
				}
			}
			out.WriteString(line + "\n")
		}
	}
	if !env.WriteFile(toFile, out.Bytes()) {
		return
	}
	env.Printf("written '%s'\n", toFile)
} //                                                                 replaceTime

// end
