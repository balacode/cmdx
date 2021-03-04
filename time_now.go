// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                       cmdx/[time_now.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

// -----------------------------------------------------------------------------
// # Command Function

// timeNow _ _
func timeNow(cmd Command, args []string) {
	tm := time.Now()
	s := tm.Format("2006-01-02 15:04:05")
	fmt.Println(s)
} //                                                                     timeNow

// end
