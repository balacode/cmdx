// -----------------------------------------------------------------------------
// CMDX Utilities Suite                               cmdx/cxfunc/[func_test.go]
// (c) balarabe@protonmail.com                                    License: GPLv3
// -----------------------------------------------------------------------------

package cxfunc

//  to test all items in func.go use:
//      go test --run Test_func_
//
//  to generate a test coverage report for the whole module use:
//      go test -coverprofile cover.out
//      go tool cover -html=cover.out

import (
	"bytes"
	"testing"
	"time"

	"github.com/balacode/zr"
)

// go test --run Test_func_ParseDuration_
func Test_func_ParseDuration_(t *testing.T) {
	zr.TBegin(t)
	// ParseDuration(s string) time.Duration
	//
	test := func(
		input string,
		wantDuration time.Duration,
		wantError bool,
	) {
		gotDuration, err := ParseDuration(input)
		gotError := err != nil
		ok1 := zr.TEqual(t, gotDuration, wantDuration)
		ok2 := zr.TEqual(t, gotError, wantError)
		if !ok1 || !ok2 {
			return // place breakpoint here
		}
	}
	const (
		HasE  = true
		NoE   = false
		Day   = time.Hour * 24
		Week  = Day * 7
		Month = Day * 30
		Year  = Day * 365
	)
	// input, wantDuration, wantError:
	//
	test("0 days", 0, NoE)
	test("0 DAYS", 0, NoE)
	test("0day", 0, NoE)
	test("0days", 0, NoE)
	test("1 DAY", Day, NoE)
	test("1 DAYS", Day, NoE)
	test("1.5day", Day/2*3, NoE)
	test("1.5days", Day/2*3, NoE)
	test("123 day", 123*Day, NoE)
	test("123 days", 123*Day, NoE)
	test("123day", 123*Day, NoE)
	test("123days", 123*Day, NoE)
	test("2day", 2*Day, NoE)
	test("2days", 2*Day, NoE)
	//
	// check unit names
	test("123 YEAR", 123*Year, NoE)
	test("234 MONTH", 234*Month, NoE)
	test("345 WEEK", 345*Week, NoE)
	test("456 DAY", 456*Day, NoE)
	test("567 HOUR", 567*time.Hour, NoE)
	test("678 MINUTE", 678*time.Minute, NoE)
	test("789 SECOND", 789*time.Second, NoE)
	//
	test("123 YEARS", 123*Year, NoE)
	test("234 MONTHS", 234*Month, NoE)
	test("345 WEEKS", 345*Week, NoE)
	test("456 DAYS", 456*Day, NoE)
	test("567 HOURS", 567*time.Hour, NoE)
	test("678 MINUTES", 678*time.Minute, NoE)
	test("789 SECONDS", 789*time.Second, NoE)
	//
	test("123 YEAR(S)", 123*Year, NoE)
	test("234 MONTH(S)", 234*Month, NoE)
	test("345 WEEK(S)", 345*Week, NoE)
	test("456 DAY(S)", 456*Day, NoE)
	test("567 HOUR(S)", 567*time.Hour, NoE)
	test("678 MINUTE(S)", 678*time.Minute, NoE)
	test("789 SECOND(S)", 789*time.Second, NoE)
	//
	// single-letter formats for years and days are not accepted
	test("123 d", 0, HasE)
	test("234 D", 0, HasE)
	test("345 w", 0, HasE)
	test("456 W", 0, HasE)
	test("567 y", 0, HasE)
	test("678 Y", 0, HasE)
	//
	// blank strings
	test("  ", 0, HasE)
	test(" ", 0, HasE)
	test("", 0, HasE)
	test("\n", 0, HasE)
	test("\r", 0, HasE)
	test("\t", 0, HasE)
	//
	// missing value
	test("day", 0, HasE)
	test("days", 0, HasE)
	test("year", 0, HasE)
	test("years", 0, HasE)
	//
	// try various input permutations
	var Spaces = []string{"", " ", "  ", "\t"}
	var Days = []string{"day", "days", "day(s)", "DAY", "DAYS", "DAY(s)"}
	for _, s := range permuteStrings(
		Spaces, []string{"", "0"}, []string{"0", "0.", "0.0"},
		Spaces, Days, Spaces,
	) {
		test(s, 0, NoE)
	}
	for _, s := range permuteStrings(
		Spaces, []string{"", "0"}, []string{"1", "1.", "1.0"},
		Spaces, Days, Spaces,
	) {
		test(s, Day, NoE)
	}
	for _, s := range permuteStrings(
		Spaces, []string{"", "0"}, []string{"1.5"},
		Spaces, Days, Spaces,
	) {
		test(s, Day/2*3, NoE)
	}
	for _, s := range permuteStrings(
		Spaces, []string{"", "0"}, []string{"2", "2.0"},
		Spaces, Days, Spaces,
	) {
		test(s, 2*Day, NoE)
	}
} //                                                    Test_func_ParseDuration_

// permuteStrings returns all combinations of strings in 'parts'
func permuteStrings(parts ...[]string) (ret []string) {
	{
		n := 1
		for _, ar := range parts {
			n *= len(ar)
		}
		ret = make([]string, 0, n)
	}
	at := make([]int, len(parts))
	var buf bytes.Buffer
loop:
	for {
		// increment position counters
		for i := len(parts) - 1; i >= 0; i-- {
			if at[i] > 0 && at[i] >= len(parts[i]) {
				if i == 0 || (i == 1 && at[i-1] == len(parts[0])-1) {
					break loop
				}
				at[i] = 0
				at[i-1]++
			}
		}
		// construct permutated string
		buf.Reset()
		for i, ar := range parts {
			j := at[i]
			if j >= 0 && j < len(ar) {
				buf.WriteString(ar[j])
			}
		}
		ret = append(ret, buf.String())
		at[len(parts)-1]++
	}
	return ret
} //                                                              permuteStrings

// end
