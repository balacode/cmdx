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

// replaceTime _ _
func replaceTime(cmd Command, args []string) {
	if len(args) < 2 {
		env.Println("'replace-time' requires: <source file> and <target file>")
		return
	}
	var (
		fromFile = args[0]
		toFile   = args[1]
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
			tm := rtExtractTime(line)
			if tm == "" {
				continue
			}
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
			tm := rtExtractTime(line)
			if tm != "" {
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
}

var (
	// `YYYY-MM-DD hh:mm ` and `YYYY-MM-DD hh:mm:ss ` (note the ending space)
	rtValidTime1 = regexp.MustCompile(`^\d{4}-\d\d-\d\d \d\d:\d\d `)
	rtValidTime2 = regexp.MustCompile(`^\d{4}-\d\d-\d\d \d\d:\d\d:\d\d `)
)

// rtExtractTime _ _
func rtExtractTime(line string) string {
	var ret string
	switch {
	case rtValidTime2.MatchString(line):
		ret = line[:19] + ":00" // 19 = length of 'YYYY-MM-DD hh:mm:ss'
	case rtValidTime1.MatchString(line):
		ret = line[:16] // 16 = length of 'YYYY-MM-DD hh:mm' in characters'
	}
	return ret
}

// end
