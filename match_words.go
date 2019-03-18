// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-03-18 01:07:59 6449F5                          cmdx/[match_words.go]
// -----------------------------------------------------------------------------

package main

import (
	str "strings"

	"github.com/balacode/zr"
)

// matchWords __
func matchWords(cmd Command, args []string) {
	if len(args) != 2 {
		env.Println("requires <word-length> and <letter-set> parameters")
		return
	}
	length := zr.Int(args[0])
	letterSet := args[1]
	// load word list
	var words []string
	{
		data, done := env.ReadFile(WordListFile)
		if !done {
			return
		}
		s := string(data)
		s = str.Replace(s, "\r"+LF, LF, -1)
		for str.Contains(s, LF+LF) {
			s = str.Replace(s, LF+LF, LF, -1)
		}
		words = str.Split(s, LF)
	}
	for _, word := range words {
		if len(word) != length {
			continue
		}
		word = str.ToLower(word)
		used := 0
		letters := str.Split(letterSet, "")
		for _, wordLetter := range str.Split(word, "") {
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
