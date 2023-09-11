package schedule

import "time"

type Position string

const (
	First  Position = "first"
	Second Position = "second"
	Third  Position = "third"
	Last   Position = "last"
	Every  Position = "every"
)

type WeekdayRule struct {
	Day   time.Weekday
	Occur Position
}

func (wr *WeekdayRule) HappensOn(date time.Time) bool {

	if date.Weekday().String() != wr.Day.String() {
		return false
	}

	switch wr.Occur {
	case Every:
		return true
	case Last:
		return date.AddDate(0, 0, 7).Month() != date.Month()
	case First:
		return date.Day() <= 7
	case Second:
		return date.Day() >= 8 && date.Day() <= 14
	case Third:
		return date.Day() >= 15 && date.Day() <= 21
	}

	return false
}
