// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                      cmdx/[time_item.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

// TimeItem _ _
type TimeItem struct {
	Time  time.Time
	Text  string
	Count int
	Spent time.Duration
}

// String _ _
func (ob *TimeItem) String() string {
	date := ob.Time.Format("2006-01-02")
	hours := ob.Spent.Hours()
	return fmt.Sprintf("%s %7.2f %11s: %s", date, hours, ob.Spent, ob.Text)
}

// -----------------------------------------------------------------------------

// TimeItemsByTime _ _
// use sort.Sort(TimeItemsByTime(ar)) to sort
type TimeItemsByTime []TimeItem

// Len _ _
func (a TimeItemsByTime) Len() int { return len(a) }

// Swap _ _
func (a TimeItemsByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less _ _
func (a TimeItemsByTime) Less(i, j int) bool {
	return a[i].Time.Sub(a[j].Time) < 0
}

// -----------------------------------------------------------------------------

// TimeItemsByDateAndDescSpent _ _
// use sort.Sort(TimeItemsByDateAndDescSpent(ar)) to sort
type TimeItemsByDateAndDescSpent []TimeItem

// Len _ _
func (a TimeItemsByDateAndDescSpent) Len() int { return len(a) }

// Swap _ _
func (a TimeItemsByDateAndDescSpent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less _ _
func (a TimeItemsByDateAndDescSpent) Less(i, j int) bool {
	if a[i].Time.Sub(a[j].Time) < 0 {
		return true
	}
	if a[i].Time == a[j].Time {
		return a[i].Spent > a[j].Spent // descending
	}
	return false
}

// end
