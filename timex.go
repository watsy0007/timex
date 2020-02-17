package timex

import "time"

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

var (
	DaysMonth     = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	DaysMonthLeap = []int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	NSec          = int(time.Second - time.Nanosecond)
)

func AtBeginningOfDay(target time.Time) time.Time {
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, target.Location())
}

func AtBeginningOfHour(target time.Time) time.Time {
	return time.Date(target.Year(), target.Month(), target.Day(), target.Hour(), 0, 0, 0, target.Location())
}

func AtBeginningOfMinute(target time.Time) time.Time {
	return time.Date(target.Year(), target.Month(), target.Day(), target.Hour(), target.Minute(), 0, 0, target.Location())
}

func AtBeginningOfMonth(target time.Time) time.Time {
	return time.Date(target.Year(), target.Month(), 1, 0, 0, 0, 0, target.Location())
}

func AtBeginningOfYear(target time.Time) time.Time {
	return time.Date(target.Year(), 1, 1, 0, 0, 0, 0, target.Location())
}

func AtBeginningOfQuarter(target time.Time) time.Time {
	month := (target.Month()-1/3)*3 + 1
	return time.Date(target.Year(), month, 1, 0, 0, 0, 0, target.Location())
}

func AtBeginningOfSemester(target time.Time) time.Time {
	month := (target.Month()-1/6)*6 + 1
	return time.Date(target.Year(), month, 1, 0, 0, 0, 0, target.Location())
}

func AtBeginningOfWeek(target time.Time) time.Time {
	dow := target.Weekday()
	if dow == time.Sunday {
		return AtBeginningOfDay(target).Add(-6 * Day)
	}
	return AtBeginningOfDay(target).AddDate(0, 0, -int(dow-time.Sunday)-1)
}

func AtEndOfHour(target time.Time) time.Time {
	y, m := target.Year(), target.Month()
	return time.Date(y, m, target.Day(), target.Hour(), 59, 59, NSec, target.Location())
}

func AtEndOfMinute(target time.Time) time.Time {
	y, m := target.Year(), target.Month()
	return time.Date(y, m, target.Day(), target.Hour(), target.Minute(), 59, NSec, target.Location())
}

func AtEndOfQuarter(target time.Time) time.Time {
	y, m := target.Year(), target.Month()
	var month time.Month
	var day int
	if m <= time.June {
		month, day = time.June, 30
	} else {
		month, day = time.December, 31
	}
	return time.Date(y, month, day, 23, 59, 59, NSec, target.Location())
}

func AtEndOfSemester(target time.Time) time.Time {
	y, m := target.Year(), target.Month()
	var month time.Month
	var day int
	if m <= time.March {
		month, day = time.March, 31
	} else if m <= time.June {
		month, day = time.June, 30
	} else if m <= time.September {
		month, day = time.September, 30
	} else {
		month, day = time.December, 31
	}
	return time.Date(y, month, day, 23, 59, 59, NSec, target.Location())
}

func AtEndOfWeek(target time.Time) time.Time {
	dow := target.Weekday()
	if dow == time.Sunday {
		return AtEndOfDay(target)
	}
	return AtEndOfDay(target).AddDate(0, 0, 7-int(dow-time.Sunday))
}

func AtEndOfDay(target time.Time) time.Time {
	y, m := target.Year(), target.Month()
	return time.Date(y, m, target.Day(), 23, 59, 59, NSec, target.Location())
}

func AtEndOfMonth(target time.Time) time.Time {
	y, m := target.Year(), target.Month()
	return time.Date(y, m, daysInMonth(y, int(m)), 23, 59, 59, NSec, target.Location())
}

func AtEndOfYear(target time.Time) time.Time {
	return time.Date(target.Year(), 12, 31, 23, 59, 59, NSec, target.Location())
}

func AtMidday(target time.Time) time.Time {
	y, m, d := target.Year(), target.Month(), target.Day()
	return time.Date(y, m, d, 12, 0, 0, 0, target.Location())
}

func daysInMonth(year, month int) int {
	if month < 1 || month > 12 {
		panic("invalid month")
	}
	if year < 1 || year > 99999 {
		panic("invalid year")
	}
	if IsLeapYear(year) {
		return DaysMonthLeap[month]
	}
	return DaysMonth[month]
}

func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
