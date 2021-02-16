// -----------------------------------------------------------------------------
// CMDX Utility                                            cmdx/cxfunc/[func.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package cxfunc

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ParseDuration parses a duration string, similar to time.ParseDuration,
// but it also processes additional long-format time durations:
//
// "year", with a year treated as 365 days
//
// "month", with a month treated as 30 days
//
// "week", "day", "hour", "minute", "second"
//
// You can also specify plural forms like "years" or "year(s)".
//
// However, this function does not allow combination of years, months and
// weeks with each other or with smaller units in the same input string,
// so a string like "1 year 7 months" will not be parsed. Instead you
// could use a fraction like "1.583y".
//
// From time.ParseDuration:
//
// A duration string is a possibly signed sequence of
// decimal numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m".
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
//
func ParseDuration(s string) (time.Duration, error) {
	if dur, err := time.ParseDuration(s); err == nil {
		return dur, nil
	}
	s = strings.ToLower(strings.TrimSpace(s))
	unit := time.Duration(0)
	unitStr := strings.TrimSpace(strings.Trim(s, "-0123456789."))
	if strings.HasSuffix(unitStr, "(s)") {
		unitStr = unitStr[:len(unitStr)-3]
	} else if strings.HasSuffix(unitStr, "s") {
		unitStr = unitStr[:len(unitStr)-1]
	}
	const Day = 24 * time.Hour
	switch unitStr {
	case "year":
		unit = Day * 365
	case "month":
		unit = Day * 30
	case "week":
		unit = Day * 7
	case "day":
		unit = Day
	case "hour":
		unit = time.Hour
	case "minute":
		unit = time.Minute
	case "second":
		unit = time.Second
	}
	if unit > 0 {
		if i := strings.Index(s, unitStr); i > 0 {
			s = strings.TrimSpace(s[:i])
			if n, err := strconv.ParseFloat(s, 32); err == nil {
				dur := time.Duration(n * float64(unit))
				return dur, nil
			}
		}
	}
	errorMsg := "invalid duration \"" + s + "\""
	return 0, errors.New(errorMsg)
} //                                                               ParseDuration

// end
