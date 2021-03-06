//
//    Description:  You are given the following information, but you may prefer
//                  to do some research for yourself.
//                    - 1 Jan 1900 was a Monday.
//                    - Thirty days has September,
//                      April, June and November.
//                      All the rest have thirty-one,
//                      Saving February alone,
//                      Which has twenty-eight, rain or shine.
//                      And on leap years, twenty-nine.
//                    - A leap year occurs on any year evenly divisible by 4,
//                      but not on a century unless it is divisible by 400.
//
//       Question:  How many Sundays fell on the first of the month during the
//                  twentieth century (1 Jan 1901 to 31 Dec 2000)?
//
//       Compiler:  go
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

//------------------------------------------------------------------------------
//  Includes
//------------------------------------------------------------------------------
import (
	"fmt"
)

//------------------------------------------------------------------------------
//  Types
//------------------------------------------------------------------------------
type Year int
type Month int
type Day int

//------------------------------------------------------------------------------
//  Enumerates
//------------------------------------------------------------------------------
// Weekdays
const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// Months
const (
	//_ = iota
	Jan = iota
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
	NrOfMonths
)

func main() {
	fmt.Println("Number of Sundays", GetNumberOfSundays(1901, 2001))
}

//-----------------------------------------------------------------------------
//  Main function implementations
//-----------------------------------------------------------------------------
// {{{
// ===  FUNCTION  ==============================================================
//         Name:  GetNumberOfSundays
//  Description:  Calculates number of sundays
// =============================================================================
func GetNumberOfSundays(startYear Year, endYear Year) Day {

	// Since 1 Jan 1900 was a Monday, we calculate the offset of days since the turn
	// of the century
	days := GetDaysOfYear(1900)
	var nrOfSundays Day

	for year := startYear; year != endYear; year++ {
		var month Month
		for month = 0; month != NrOfMonths; month++ {
			var daysOfThisMonth Day
			for daysOfThisMonth = 1; daysOfThisMonth != GetDaysOfMonth(month, year)+1; daysOfThisMonth++ {

				days++

				if days%Sunday == 0 && daysOfThisMonth == 1 {
					nrOfSundays++
				} // if

			} // days
		} // month
	} // year

	return nrOfSundays
}

// ===  FUNCTION  ==============================================================
//         Name:  GetDaysOfYear
//  Description:  Calculates number of days for a given month and year
// =============================================================================
func GetDaysOfYear(year Year) Day {
	var days Day

	var month Month
	for month = 0; month != NrOfMonths; month++ {
		var daysOfThisMonth Day
		for daysOfThisMonth = 1; daysOfThisMonth != GetDaysOfMonth(month, year)+1; daysOfThisMonth++ {
			days++
		} // days
	} //month

	return days
}

// ===  FUNCTION  ==============================================================
//         Name:  GetDaysOfMonth
//  Description:  Calculates number of days for a given month and year
// =============================================================================
func GetDaysOfMonth(month Month, year Year) Day {
	var daysInThisMonth Day

	switch month {
	case Jan, Mar, May, Jul, Aug, Oct, Dec:
		daysInThisMonth = 31
	case Feb:
		if IsLeapYear(year) == true {
			daysInThisMonth = 29
		} else {
			daysInThisMonth = 28
		}
	case Apr, Jun, Sep, Nov:
		daysInThisMonth = 30
	default:
		daysInThisMonth = -1
	}

	return daysInThisMonth
}

// ===  FUNCTION  ==============================================================
//         Name:  IsLeapYear
//  Description:  Checks if a given year is a leap year
// =============================================================================
func IsLeapYear(year Year) bool {

	isLeapYear := false

	if year%4 == 0 {
		switch {
		case year%100 == 0 && year%400 != 0:
			isLeapYear = false
		default:
			isLeapYear = true
		}
	}

	return isLeapYear
}

//}}}
