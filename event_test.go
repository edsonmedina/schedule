package schedule

import (
	"testing"
	"time"
)

func TestEvent_HappensOn(t *testing.T) {

	parseDate := func(s string) time.Time {
		t, err := time.Parse("2006-01-02", s)
		if err != nil {
			panic(err)
		}

		return t
	}

	t.Run("Fixed date", func(t *testing.T) {
		st := &Event{
			Date: parseDate("2021-01-01"),
		}

		if !st.HappensOn(parseDate("2021-01-01")) {
			t.Errorf("Expected to happen on 2021-01-01")
		}

		if st.HappensOn(parseDate("2021-01-02")) {
			t.Errorf("Expected to not happen on 2021-01-02")
		}
	})

	t.Run("Daily", func(t *testing.T) {
		st := &Event{
			Frequency: Daily,
			StartDate: parseDate("2021-01-01"),
			EndDate:   parseDate("2021-01-31"),
		}

		if !st.HappensOn(parseDate("2021-01-01")) {
			t.Errorf("Expected to happen on 2021-01-01")
		}

		if !st.HappensOn(parseDate("2021-01-31")) {
			t.Errorf("Expected to happen on 2021-01-31")
		}

		if st.HappensOn(parseDate("2021-02-01")) {
			t.Errorf("Expected to not happen after end date on 2021-02-01")
		}
	})

	t.Run("Weekly", func(t *testing.T) {
		st := &Event{
			Frequency:   Weekly,
			WeekdayRule: WeekdayRule{Day: time.Monday, Occur: Every},
			StartDate:   parseDate("2021-01-01"),
			EndDate:     parseDate("2021-01-31"),
		}

		if !st.HappensOn(parseDate("2021-01-04")) {
			t.Errorf("Expected to happen on 2021-01-04")
		}

		if !st.HappensOn(parseDate("2021-01-25")) {
			t.Errorf("Expected to happen on 2021-01-25")
		}

		if st.HappensOn(parseDate("2021-01-02")) {
			t.Errorf("Expected to not happen on 2021-01-02")
		}
	})

	t.Run("Monthly", func(t *testing.T) {
		st := &Event{
			Frequency:  Monthly,
			DayOfMonth: 5,
			StartDate:  parseDate("2021-01-01"),
			EndDate:    parseDate("2021-12-31"),
		}

		if !st.HappensOn(parseDate("2021-01-05")) {
			t.Errorf("Expected to happen on 2021-01-05")
		}

		if st.HappensOn(parseDate("2021-01-31")) {
			t.Errorf("Expected to not happen on 2021-01-31")
		}

		if !st.HappensOn(parseDate("2021-02-05")) {
			t.Errorf("Expected to happen on 2021-02-05")
		}

		if st.HappensOn(parseDate("2022-01-05")) {
			t.Errorf("Expected to not happen on 2022-01-05")
		}
	})

	t.Run("Yearly", func(t *testing.T) {
		st := &Event{
			Frequency:  Yearly,
			DayOfMonth: 5,
			Month:      11,
			StartDate:  parseDate("2021-01-01"),
			EndDate:    parseDate("2023-12-31"),
		}

		if !st.HappensOn(parseDate("2021-11-05")) {
			t.Errorf("Expected to happen on 2021-11-05")
		}

		if !st.HappensOn(parseDate("2022-11-05")) {
			t.Errorf("Expected to happen on 2022-11-05")
		}

		if st.HappensOn(parseDate("2021-11-06")) {
			t.Errorf("Expected to not happen on 2021-11-06")
		}

		if st.HappensOn(parseDate("2024-11-05")) {
			t.Errorf("Expected to not happen after end date on 2024-11-05")
		}
	})

	t.Run("Weekdays", func(t *testing.T) {
		st := &Event{
			Frequency: Weekdays,
			StartDate: parseDate("2021-01-01"),
			EndDate:   parseDate("2021-01-31"),
		}

		if !st.HappensOn(parseDate("2021-01-04")) {
			t.Errorf("Expected to happen on 2021-01-04")
		}

		if st.HappensOn(parseDate("2021-01-09")) {
			t.Errorf("Expected to not happen on 2021-01-09")
		}

		if !st.HappensOn(parseDate("2021-01-29")) {
			t.Errorf("Expected to happen on 2021-01-29")
		}

		if st.HappensOn(parseDate("2021-01-02")) {
			t.Errorf("Expected to not happen after end date on 2021-01-02")
		}
	})

	t.Run("Weekends", func(t *testing.T) {
		st := &Event{
			Frequency: Weekends,
			StartDate: parseDate("2021-01-01"),
			EndDate:   parseDate("2021-01-31"),
		}

		if !st.HappensOn(parseDate("2021-01-02")) {
			t.Errorf("Expected to happen on 2021-01-02")
		}

		if st.HappensOn(parseDate("2021-01-04")) {
			t.Errorf("Expected to not happen on 2021-01-04")
		}

		if !st.HappensOn(parseDate("2021-01-31")) {
			t.Errorf("Expected to happen on 2021-01-31")
		}

		if st.HappensOn(parseDate("2021-02-06")) {
			t.Errorf("Expected to not happen after end date on 2021-02-06")
		}
	})
}
