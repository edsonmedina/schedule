package schedule

import "time"

type Event struct {
	Date        time.Time   // Date of the event
	Frequency   Frequency   // e.g., daily, weekly, monthly, yearly
	DayOfMonth  int         // e.g., 5 for the 5th day of the month
	WeekdayRule WeekdayRule // e.g., {Day: "Friday", Occur: "last"}
	Month       int         // e.g., 11 for November
	StartDate   time.Time   // Start date for recurrent events
	EndDate     time.Time   // End date for recurring events
}

func (e *Event) HappensOn(t time.Time) bool {

	if e.Date == t {
		return true
	}

	if e.StartDate.After(t) || e.EndDate.Before(t) {
		return false
	}

	switch e.Frequency {
	case Daily:
		return true
	case Weekly:
		return e.WeekdayRule.HappensOn(t)
	case Monthly:
		return e.DayOfMonth == t.Day()
	case Yearly:
		return e.Month == int(t.Month()) && e.DayOfMonth == t.Day()
	case Weekdays:
		return t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
	case Weekends:
		return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
	}

	return false
}
