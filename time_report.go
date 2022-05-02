// -----------------------------------------------------------------------------
// CMDX Utilities Suite                                    cmdx/[time_report.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package main

// # Command Function
//   timeReport(cmd Command, args []string)
//
// # Report Functions
//   trMonthlySummary(minDate, maxDate interface{}, files []string)
//   trSummaryByDateText(minDate, maxDate interface{}, files []string)
//
// # Functions
//   trCalcSpent(ar []TimeItem, autoTime bool) []TimeItem
//   trFilterDates(ar []TimeItem, minDate, maxDate interface{}) []TimeItem
//   trGetTimeItems(lines []string) []TimeItem
//   trIsTimeStart(s string) bool
//   trMergeFiles(filenames []string) (lines []string)
//   trPrintFaults(ar []TimeItem)
//   trPrintTimeItems(entries []TimeItem)
//   trSumByDate(items []TimeItem) (ret []TimeItem)
//   trSumByDateText(items []TimeItem) (ret []TimeItem)
//
// # Helper Functions
//   trDateStr(val time.Time) string
//   trDateTimeStr(val time.Time) string

// For debugging, you can set "args:" in launch.json. E.g:
// "args": [
//     "time-report",
//     "2022-01-01",
//     "2022-31-12",
//     "=admin/",
//     "/x/user/admin/timelog.txt",
// ]

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/balacode/zr"
	fs "github.com/balacode/zr-fs"
)

const (
	trAuto   = 1
	trManual = 2
)

// -----------------------------------------------------------------------------
// # Command Function

// timeReport _ _
func timeReport(cmd Command, args []string) {
	var (
		min      = "1900-01-01"
		max      = "9999-12-31"
		files    = []string{}
		contains = []string{}
		dates    = 0
		dateRX   = regexp.MustCompile(`^\d{4}-\d\d-\d\d$`)
	)
	for _, arg := range args {
		switch {
		case dateRX.MatchString(arg):
			dates++
			switch dates {
			case 1:
				min = arg
			case 2:
				max = arg
			default:
				fmt.Printf("Error: too many dates")
				return
			}
		case strings.HasPrefix(arg, "="):
			contains = append(contains, arg[1:])
		case fs.FileExists(arg):
			files = append(files, arg)
		default:
			fmt.Printf("Error: file %q doesn't exist\n", arg)
			return
		}
	}
	if len(files) < 1 {
		files = []string{"timelog.txt"}
	}
	trMonthlySummary("MANUAL", trManual, min, max, files, contains)
}

// -----------------------------------------------------------------------------
// # Report Functions

// trMonthlySummary _ _
func trMonthlySummary(
	caption string,
	mode int,
	minDate, maxDate interface{},
	files []string,
	contains []string,
) {
	var (
		lenLines = 0
		cal      zr.Calendar
	)
	{
		lines := trMergeFiles(files, contains)
		lenLines = len(lines)
		items := trFilterDates(trGetTimeItems(lines), minDate, maxDate)
		if mode == trAuto {
			items = trCalcSpent(items, true)
		} else if mode == trManual {
			trPrintFaults(items)
			items = trCalcSpent(items, false)
		}
		sums := trSumByDate(items)
		cal.SetWeekTotals(true)
		for _, itm := range sums {
			if itm.Spent < 0 {
				cal.Set(itm.Time, "** >24")
			} else {
				cal.Set(itm.Time, itm.Spent.Hours())
			}
		}
	}
	if caption != "" {
		env.Println(caption)
	}
	env.Println("FROM:", minDate, "TO:", maxDate, "LINES:", lenLines)
	env.Println(cal.String())
	env.Println()
}

/*UNUSED:
// trSummaryByDateText _ _
func trSummaryByDateText(minDate, maxDate interface{}, files []string) {
	var items []TimeItem
	{
		lines := trMergeFiles(files)
		items = trGetTimeItems(lines)
		items = trFilterDates(items, minDate, maxDate)
		//
		env.Println("FROM:", minDate, "TO:", maxDate, "LINES:", len(lines))
	}
	items = trCalcSpent(items, true)
	if true {
		for _, itm := range items {
			env.Println(trDateTimeStr(itm.Time), "->", itm.String())
		}
	}
	sum := trSumByDateText(items)
	sort.Sort(TimeItemsByDateAndDescSpent(sum))
	trPrintTimeItems(sum)
}
:UNUSED*/

// -----------------------------------------------------------------------------
// # Functions

// trCalcSpent reads time entries from 'ar' and returns a
// copy with all time spent (Spent field) recalculated.
func trCalcSpent(ar []TimeItem, autoTime bool) []TimeItem {
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
		if trIsTimeStart(itm.Text) {
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
}

// trFilterDates filters times in 'ar' and returns a new slice
// of times that fall within the specified date range.
func trFilterDates(ar []TimeItem, minDate, maxDate interface{}) []TimeItem {
	//
	min := parseTime(minDate).String()[:10]
	max := parseTime(maxDate).String()[:10]
	var ret []TimeItem
	for _, t := range ar {
		date := trDateStr(t.Time)
		if date >= min && date <= max {
			ret = append(ret, t)
		}
	}
	return ret
}

// trGetTimeItems returns a TimeItem array from a string slice
// - entries in the result are in ascending order
// - each date+time is unique
//
func trGetTimeItems(lines []string) []TimeItem {
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
			Time:  parseTime(key),
			Text:  m[key],
			Count: 1,
			Spent: 0,
		})
	}
	return ret
}

// trIsTimeStart _ _
func trIsTimeStart(s string) bool {
	if strings.TrimSpace(s) == "IN" ||
		strings.HasPrefix(s, "IN ") ||
		strings.Contains(s, " IN ") ||
		strings.HasSuffix(s, " IN") {
		return true
	}
	return false
}

// trMergeFiles reads and filters lines from the specified filenames and
// returns the combined lines as a slice of strings. If 'contains' is
// specified, only returns lines that contain the string(s) it contains.
// If 'contains' is zero-length, returns all lines.
func trMergeFiles(filenames, contains []string) (lines []string) {
	for _, path := range filenames {
		fileLines := env.ReadFileLines(path)
		if len(contains) == 0 {
			lines = append(lines, fileLines...)
			continue
		}
		for _, ln := range fileLines {
		nextLine:
			for _, sub := range contains {
				if strings.Contains(strings.ToLower(ln), strings.ToLower(sub)) {
					lines = append(lines, ln)
					continue nextLine
				}
			}
		}
	}
	return lines
}

// trPrintFaults _ _
func trPrintFaults(ar []TimeItem) {
	//
	var prevDate string    // previous date
	var prevTime time.Time // previous date and time
	//
	for _, itm := range ar {
		date := trDateStr(itm.Time)
		if trIsTimeStart(itm.Text) {
			prevDate = date
			prevTime = itm.Time
			continue
		}
		if date != prevDate {
			env.Println("NO START:", trDateTimeStr(itm.Time), itm.Text)
		}
		var spent time.Duration
		if !prevTime.IsZero() {
			spent = itm.Time.Sub(prevTime)
		}
		if spent > 3*time.Hour {
			s := fmt.Sprintf("(%s)", spent)
			if !strings.Contains(itm.Text, s) {
				env.Println("TOO LONG:", trDateTimeStr(itm.Time), itm.Text, s)
			}
		}
		prevDate = date
		prevTime = itm.Time
	}
}

/*UNUSED:
// trPrintTimeItems _ _
func trPrintTimeItems(entries []TimeItem) {
	var (
		prev  string
		total time.Duration
		grand time.Duration
	)
	prt := func(a ...interface{}) {
		env.Printf("%s %7.2f %11s: %s\n", a...)
	}
	for _, t := range entries {
		date := trDateStr(t.Time)
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
}
:UNUSED*/

// trSumByCategory _ _
func trSumByCategory(items []TimeItem) (ret []TimeItem) {
	m := map[string]*TimeItem{}
	for _, t := range items {
		k := t.Text
		i := strings.Index(k, ":")
		if i > -1 {
			k = k[:i]
		}
		if _, exist := m[k]; !exist {
			m[k] = &TimeItem{Text: k}
		}
		m[k].Count += t.Count
		m[k].Spent += t.Spent
	}
	// convert items map into slice
	for _, t := range m {
		ret = append(ret, *t)
	}
	return ret
}

// trSumByDate _ _
func trSumByDate(items []TimeItem) (ret []TimeItem) {
	//
	// grouping stage
	m := map[string]*TimeItem{}
	for _, t := range items {
		dt := trDateStr(t.Time)
		if _, exist := m[dt]; !exist {
			m[dt] = &TimeItem{Time: parseTime(dt)}
		}
		if m[dt].Spent == -1 {
			continue
		}
		m[dt].Count += t.Count
		sum := m[dt].Spent + t.Spent
		if t.Spent > time.Hour*24 {
			env.Println("> 24 HRS:", dt, t.Spent)
			m[dt].Spent = -1
		} else {
			m[dt].Spent = sum
		}
	}
	// convert items map into slice
	for _, t := range m {
		ret = append(ret, *t)
	}
	return ret
}

/*UNUSED:
// trSumByDateText _ _
func trSumByDateText(items []TimeItem) (ret []TimeItem) {
	//
	// grouping stage
	m := map[string]*TimeItem{}
	for _, t := range items {
		date := trDateStr(t.Time)
		k := date + "\t" + t.Text
		if _, exist := m[k]; !exist {
			m[k] = &TimeItem{
				Time: parseTime(date),
				Text: t.Text,
			}
		}
		m[k].Count += t.Count
		m[k].Spent += t.Spent
	}
	// convert items map into slice
	for _, t := range m {
		ret = append(ret, *t)
	}
	return ret
}
:UNUSED*/

// -----------------------------------------------------------------------------
// # Helper Functions

// trDateStr returns a date-only string ("2006-01-02" format) from a Time value.
func trDateStr(val time.Time) string {
	return val.Format("2006-01-02")
}

// trDateTimeStr returns a date-time string
// ("2006-01-02 15:04:05" format) from a Time value.
func trDateTimeStr(val time.Time) string {
	return val.Format("2006-01-02 15:04:05")
}

// end
