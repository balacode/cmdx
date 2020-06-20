// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                    License: GPLv3
// :v: 2020-06-20 09:58:17 5FA480                          cmdx/[time_report.go]
// -----------------------------------------------------------------------------

// WORK-IN-PROGRESS: @2018-02-26 15:47

package main

// # Command Function
//   timeReport(cmd Command, args []string)
//
// # Report Functions
//   trptMonthlySummary(minDate, maxDate interface{}, files []string)
//   trptSummaryByDateText(minDate, maxDate interface{}, files []string)
//
// # Functions
//   trptCalcSpent(ar []TimeItem, autoTime bool) []TimeItem
//   trptFilterDates(ar []TimeItem, minDate, maxDate interface{}) []TimeItem
//   trptGetTimeItems(lines []string) []TimeItem
//   trptIsTimeStart(s string) bool
//   trptMergeFiles(filenames []string) (lines []string)
//   trptPrintFaults(ar []TimeItem)
//   trptPrintTimeItems(entries []TimeItem)
//   trptSumByDate(items []TimeItem) (ret []TimeItem)
//   trptSumByDateText(items []TimeItem) (ret []TimeItem)
//   trptSummaryByDateText(minDate, maxDate interface{}, files []string)
//
// # Helper Functions
//   dateStr(val time.Time) string
//   dateTimeStr(val time.Time) string
//   timeOf(value interface{}) time.Time

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/balacode/zr"
)

type trptMode bool

const (
	trptAuto   = 1
	trptManual = 2
)

var trptAutoLogFiles = hardcodedAutologFiles

var trptManualLogFiles = hardcodedManualLogFiles

var trptIgnoreProjects = []string{
	"cmdx",
	"demo",
	"dmd_app",
	"lsrv",
	"priveda",
	"tex2",
	"tlg",
	"whois",
}

// -----------------------------------------------------------------------------
// # Command Function

// timeReport _ _
func timeReport(cmd Command, args []string) {
	var min, max string
	/*
		min, max = "2017-11-24", "2017-11-30"
		trptMonthlySummary(min, max, trptAutoLogFiles)
		trptMonthlySummary(min, max, trptManualLogFiles)
		//
		min, max = "2017-12-01", "2017-12-31"
		trptMonthlySummary(min, max, trptAutoLogFiles)
		trptMonthlySummary(min, max, trptManualLogFiles)
		//
		min, max = "2018-01-01", "2018-01-31"
		trptMonthlySummary(min, max, trptAutoLogFiles)
		trptMonthlySummary(min, max, trptManualLogFiles)
	*/
	min, max = "2017-11-24", "2018-02-18"
	//``
	/*
		min, max = "2018-01-13", "2018-01-14"
	*/
	/*
		trptMonthlySummary("AUTO", trptAuto, min, max, trptAutoLogFiles)
	*/
	trptMonthlySummary("MANUAL", trptManual, min, max, trptManualLogFiles)
	return
	//
	min = dateStr(time.Now().Add(-24 * time.Hour))
	max = dateStr(time.Now())
	if len(args) == 1 {
		if strings.ToLower(args[0]) == "all" {
			min = "2000-01-01"
		} else {
			min = dateStr(timeOf(args[0]))
		}
	}
	if len(args) == 2 {
		min = dateStr(timeOf(args[0]))
		max = dateStr(timeOf(args[1]))
	}
	//trptSummaryByDateText(min, max, trptAutoLogFiles)
} //                                                                  timeReport

// -----------------------------------------------------------------------------
// # Report Functions

// trptMonthlySummary _ _
func trptMonthlySummary(
	caption string,
	mode int,
	minDate, maxDate interface{},
	files []string,
) {
	lines := trptMergeFiles(files)
	var items []TimeItem
	items = trptGetTimeItems(lines)
	items = trptFilterDates(items, minDate, maxDate)
	if mode == trptAuto {
		items = trptCalcSpent(items, true)
	} else if mode == trptManual {
		trptPrintFaults(items)
		items = trptCalcSpent(items, false)
	}
	sums := trptSumByDate(items)
	env.Println(caption)
	var cal zr.Calendar
	for _, itm := range sums {
		cal.Set(itm.Time, itm.Spent.Hours())
	}
	env.Println("FROM:", minDate, "TO:", maxDate, "LINES:", len(lines))
	env.Println(cal.String())
	env.Println()
} //                                                          trptMonthlySummary

// trptSummaryByDateText _ _
func trptSummaryByDateText(minDate, maxDate interface{}, files []string) {
	var items []TimeItem
	{
		lines := trptMergeFiles(files)
		items = trptGetTimeItems(lines)
		items = trptFilterDates(items, minDate, maxDate)
		//
		env.Println("FROM:", minDate, "TO:", maxDate, "LINES:", len(lines))
	}
	items = trptCalcSpent(items, true)
	if true {
		for _, itm := range items {
			env.Println(dateTimeStr(itm.Time), "->", itm.String())
		}
	}
	sum := trptSumByDateText(items)
	sort.Sort(TimeItemsByDateAndDescSpent(sum))
	trptPrintTimeItems(sum)
} //                                                       trptSummaryByDateText

// -----------------------------------------------------------------------------
// # Functions

// trptCalcSpent reads time entries from 'ar' and returns a
// copy with all time spent (Spent field) recalculated
func trptCalcSpent(ar []TimeItem, autoTime bool) []TimeItem {
	var prev time.Time // previous time
	var ret []TimeItem
	//
	// calculate total time spent on each item
	for _, itm := range ar {
		var spent time.Duration
		if !prev.IsZero() {
			spent = itm.Time.Sub(prev)
		}
		if autoTime && spent > 10*time.Minute {
			spent = 0
		}
		if trptIsTimeStart(itm.Text) {
			spent = 0
		}
		prev = itm.Time
		ret = append(ret, TimeItem{
			Time:  itm.Time,
			Text:  itm.Text,
			Count: itm.Count,
			Spent: spent,
		})
	}
	return ret
} //                                                               trptCalcSpent

// trptFilterDates _ _
func trptFilterDates(ar []TimeItem, minDate, maxDate interface{}) []TimeItem {
	//
	min := timeOf(minDate).String()[:10]
	max := timeOf(maxDate).String()[:10]
	var ret []TimeItem
	for _, t := range ar {
		date := dateStr(t.Time)
		if date >= min && date <= max {
			ret = append(ret, t)
		}
	}
	return ret
} //                                                             trptFilterDates

// trptGetTimeItems returns a TimeItem array from a string slice
// - entries in the result are in ascending order
// - each date+time is unique
func trptGetTimeItems(lines []string) []TimeItem {
	//
	// read lines into a unique date+time map
	m := map[string]string{}
	for _, line := range lines {
		//
		// ignore lines without date/time
		if len(line) < 20 {
			continue
		}
		dt, err := time.Parse("2006-01-02 15:04:05", line[:19])
		if err != nil || dt.IsZero() {
			continue
		}
		// store each line in a unique date+time key
		// (any previous entry with same date+time gets overwritten)
		m[line[:19]] = strings.Trim(line[20:], SPACES+"/\\")
	}
	// create a sorted array of keys
	times := make([]string, 0, len(m))
	for key := range m {
		times = append(times, key)
	}
	times = sortUniqueStrings(times)
	//
	var ret []TimeItem
	for _, key := range times {
		ret = append(ret, TimeItem{
			Time:  timeOf(key),
			Text:  m[key],
			Count: 1,
			Spent: 0,
		})
	}
	return ret
} //                                                            trptGetTimeItems

// trptIsTimeStart _ _
func trptIsTimeStart(s string) bool {
	if strings.HasPrefix(s, "IN ") ||
		strings.Contains(s, " IN ") ||
		strings.HasSuffix(s, " IN") {
		return true
	}
	return false
} //                                                             trptIsTimeStart

// trptMergeFiles _ _
func trptMergeFiles(filenames []string) (lines []string) {
	for _, path := range filenames {
		lines = append(lines, env.ReadFileLines(path)...)
	}
	return lines
} //                                                              trptMergeFiles

// trptPrintFaults _ _
func trptPrintFaults(ar []TimeItem) {
	//
	var pdate string    // previous date
	var ptime time.Time // previous date and time
	//
	for _, itm := range ar {
		date := dateStr(itm.Time)
		if trptIsTimeStart(itm.Text) {
			pdate = date
			ptime = itm.Time
			continue
		}
		if date != pdate {
			env.Println("NO START:", dateTimeStr(itm.Time), itm.Text)
		}
		var spent time.Duration
		if !ptime.IsZero() {
			spent = itm.Time.Sub(ptime)
		}
		if spent > 3*time.Hour {
			s := fmt.Sprintf("(%s)", spent)
			if !strings.Contains(itm.Text, s) {
				env.Println("TOO LONG:", dateTimeStr(itm.Time), itm.Text, s)
			}
		}
		pdate = date
		ptime = itm.Time
	}
} //                                                             trptPrintFaults

// trptPrintTimeItems _ _
func trptPrintTimeItems(entries []TimeItem) {
	var (
		prev  string
		total time.Duration
		grand time.Duration
	)
	prt := func(a ...interface{}) {
		env.Printf("%s %7.2f %11s: %s\n", a...)
	}
	for _, t := range entries {
		date := dateStr(t.Time)
		if date != prev {
			if prev != "" {
				env.Println(strings.Repeat("-", 35))
				prt(prev, total.Hours(), total, "total")
			}
			env.Println()
			prev = date
			total = 0
		}
		prt(date, t.Spent.Hours(), t.Spent, t.Text)
		grand += t.Spent
		total += t.Spent
	}
	if prev != "" {
		env.Println(strings.Repeat("-", 35))
		prt(prev, total.Hours(), total, "total")
	}
	env.Println()
	env.Println(strings.Repeat("=", 35))
	prt(prev, grand.Hours(), grand, "grand total")
} //                                                          trptPrintTimeItems

// trptSumByDate _ _
func trptSumByDate(items []TimeItem) (ret []TimeItem) {
	//
	// grouping stage
	m := map[string]*TimeItem{}
	for _, t := range items {
		dt := dateStr(t.Time)
		if _, exist := m[dt]; !exist {
			m[dt] = &TimeItem{Time: timeOf(dt)}
		}
		m[dt].Count += t.Count
		m[dt].Spent += t.Spent
		//PL(dateTimeStr(t.Time), t.String(), "---->", m[dt].Spent) //``
	}
	// extract items in map into slice
	for _, t := range m {
		ret = append(ret, *t)
	}
	return ret
} //                                                               trptSumByDate

// trptSumByDateText _ _
func trptSumByDateText(items []TimeItem) (ret []TimeItem) {
	//
	// grouping stage
	m := map[string]*TimeItem{}
	for _, t := range items {
		date := dateStr(t.Time)
		k := date + "\t" + t.Text
		if _, exist := m[k]; !exist {
			m[k] = &TimeItem{
				Time: timeOf(date),
				Text: t.Text,
			}
		}
		m[k].Count += t.Count
		m[k].Spent += t.Spent
	}
	// extract items in map into slice
	for _, t := range m {
		ret = append(ret, *t)
	}
	return ret
} //                                                           trptSumByDateText

// -----------------------------------------------------------------------------
// # Helper Functions

// dateStr _ _
func dateStr(val time.Time) string {
	return val.Format("2006-01-02")
} //                                                                     dateStr

// dateTimeStr _ _
func dateTimeStr(val time.Time) string {
	return val.Format("2006-01-02 15:04:05")
} //                                                                 dateTimeStr

// timeOf converts any string-like value to time.Time without returning
// an error if the conversion failed, in which case it logs an error
// and returns a zero-value time.Time.
//
// If val is a zero-length string, returns a zero-value time.Time
// but does not log a warning.
//
// It also accepts a time.Time as input.
//
// In both cases the returned Time type will contain only the date
// part without the time or time zone components.
//
// Note: fmt.Stringer (or fmt.GoStringer) interfaces are not treated as
// strings to avoid bugs from implicit conversion. Use the String method.
//
func timeOf(value interface{}) time.Time {
	switch v := value.(type) {
	case time.Time:
		{
			return v
		}
	case string:
		{
			if v == "" {
				return time.Time{}
			}
			var tm time.Time
			var err error
			if len(v) == 10 {
				tm, err = time.Parse("2006-01-02", v)
				if err == nil && !tm.IsZero() {
					return timeOf(tm)
				}
			}
			if len(v) == 19 {
				tm, err = time.Parse("2006-01-02 15:04:05", v)
				if err == nil && !tm.IsZero() {
					return timeOf(tm)
				}
			}
			if err != nil {
				zr.Error(err)
			}
			return time.Time{}
		}
	case *string:
		if v != nil {
			return timeOf(*v)
		}
	}
	zr.Error("Can not convert", reflect.TypeOf(value), "to int:", value)
	return time.Time{}
} //                                                                      timeOf

//end
