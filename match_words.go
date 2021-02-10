// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2021-02-10 09:35:34 610EBB                          cmdx/[match_words.go]
// -----------------------------------------------------------------------------

package main

import (
	"strings"

	"github.com/balacode/zr"
)

// matchWords _ _
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
		s = strings.ReplaceAll(s, "\r\n", "\n")
		for strings.Contains(s, "\n\n") {
			s = strings.ReplaceAll(s, "\n\n", "\n")
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

// end
