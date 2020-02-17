package timex

import "time"

var local = func() time.Time { return time.Now().Local() }

func CurrentEndOfWeek() time.Time {
	return AtEndOfWeek(local())
}

func NextEndOfWeek() time.Time {
	return AtEndOfWeek(local()).AddDate(0, 0, 7)
}

func CurrentBeginningOfWeek() time.Time {
	return AtBeginningOfWeek(local())
}

func NextBeginningOfWeek() time.Time {
	return AtBeginningOfWeek(local()).AddDate(0, 0, 7)
}

func CurrentEndOfQuarter() time.Time {
	return AtEndOfQuarter(local())
}

func NextEndOfQuarter() time.Time {
	next := AtEndOfQuarter(local()).AddDate(0, 0, 1)
	return AtEndOfQuarter(next)
}

func CurrentBeginningOfQuarter() time.Time {
	return AtBeginningOfQuarter(local())
}

func NextBeginningOfQuarter() time.Time {
	next := AtEndOfQuarter(local()).AddDate(0, 0, 1)
	return AtBeginningOfQuarter(next)
}
