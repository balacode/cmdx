// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-03-01 16:49:00 DB3124                          [cmdx/match_words.go]
// -----------------------------------------------------------------------------

package main

import "strings" // standard

import "github.com/balacode/zr" // Zircon-Go

// matchWords __
func matchWords(cmd Command, args []string) {
	if len(args) != 2 {
		env.Println("requires <word-length> and <letter-set> parameters")
		return
	}
	var length = zr.Int(args[0])
	var letterSet = args[1]
	// load word list
	var words []string
	{
		var data, done = env.ReadFile(WordListFile)
		if !done {
			return
		}
		var s = string(data)
		s = strings.Replace(s, "\r"+LF, LF, -1)
		for strings.Contains(s, LF+LF) {
			s = strings.Replace(s, LF+LF, LF, -1)
		}
		words = strings.Split(s, LF)
	}
	for _, word := range words {
		if len(word) != length {
			continue
		}
		word = strings.ToLower(word)
		var used = 0
		var letters = strings.Split(letterSet, "")
		for _, wordLetter := range strings.Split(word, "") {
			for i, letter := range letters {
				if wordLetter == letter {
					used++
					letters[i] = "0"
					break
				}
			}
			if used == length {
				env.Printf("    %s    ", word)
			}
		}
	}
} //                                                              matchWords

//end
