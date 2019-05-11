// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2019-05-11 04:25:01 07A48A                          cmdx/[match_words.go]
// -----------------------------------------------------------------------------

package main

import (
	"strings"

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
		s = strings.Replace(s, "\r\n", "\n", -1)
		for strings.Contains(s, "\n\n") {
			s = strings.Replace(s, "\n\n", "\n", -1)
		}
		words = strings.Split(s, "\n")
	}
	for _, word := range words {
		if len(word) != length {
			continue
		}
		word = strings.ToLower(word)
		var (
			used    = 0
			letters = strings.Split(letterSet, "")
		)
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
} //                                                                  matchWords

//end
