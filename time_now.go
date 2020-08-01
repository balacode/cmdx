// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2020-08-01 22:30:13 9783A8                             cmdx/[time_now.go]
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
} //                                                                  timeReport

//end
