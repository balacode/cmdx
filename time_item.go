// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2018-05-09 01:03:17 3D488E                            cmdx/[time_item.go]
// -----------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"
)

// TimeItem __
type TimeItem struct {
	Time  time.Time
	Text  string
	Count int
	Spent time.Duration
}

// String __
func (ob *TimeItem) String() string {
	var date = ob.Time.Format("2006-01-02")
	var hours = ob.Spent.Hours()
	return fmt.Sprintf("%s %7.2f %11s: %s", date, hours, ob.Spent, ob.Text)
} //                                                                      String

// -----------------------------------------------------------------------------

// TimeItemsByTime __
// use sort.Sort(TimeItemsByTime(ar)) to sort
type TimeItemsByTime []TimeItem

// Len __
func (a TimeItemsByTime) Len() int { return len(a) }

// Swap __
func (a TimeItemsByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less __
func (a TimeItemsByTime) Less(i, j int) bool {
	return a[i].Time.Sub(a[j].Time) < 0
} //                                                                        Less

// -----------------------------------------------------------------------------

// TimeItemsByDateAndDescSpent __
// use sort.Sort(TimeItemsByDateAndDescSpent(ar)) to sort
type TimeItemsByDateAndDescSpent []TimeItem

// Len __
func (a TimeItemsByDateAndDescSpent) Len() int { return len(a) }

// Swap __
func (a TimeItemsByDateAndDescSpent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less __
func (a TimeItemsByDateAndDescSpent) Less(i, j int) bool {
	if a[i].Time.Sub(a[j].Time) < 0 {
		return true
	}
	if a[i].Time == a[j].Time {
		return a[i].Spent > a[j].Spent // descending
	}
	return false
} //                                                                        Less

//end
