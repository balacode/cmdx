// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2021-02-10 09:33:24 69460C                             cmdx/[time_now.go]
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
