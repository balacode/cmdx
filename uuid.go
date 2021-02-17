// -----------------------------------------------------------------------------
// CMDX Utility                                                   cmdx/[uuid.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
)

// printUUID generates and prints one or mode UUIDs,
// also known as Universally Unique Identifiers.
// The format is 'XXXXXXXX-XXXX-4XXX-ZXXX-XXXXXXXXXXXX' where every X is a
// random upper-case hex digit, and Z must be one of '8', '9', 'A' or 'B'.
func printUUID(cmd Command, args []string) {
	count := 0
	for _, arg := range args {
		n, _ := strconv.ParseInt(arg, 10, 32) // base 10, bitSize 32
		count += int(n)
	}
	if count == 0 {
		count = 1
	}
	for count > 0 {
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			continue
		}
		// 13th character (at [12]) must be '4'
		b[6] = (b[6] | 0x40) & 0x4F
		//
		// 17th character at [16] must be '8', '9', 'A', or 'B'
		b[8] = (b[8] | 0x80) & 0xBF
		//
		s := fmt.Sprintf("%X-%X-%X-%X-%X",
			b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
		fmt.Println(s)
		count--
	}
} //                                                                   printUUID

// end
