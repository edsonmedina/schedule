package schedule

import (
	"slices"
	"testing"
	"time"
)

func TestListEventsBetween(t *testing.T) {
	start := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC)

	event1 := Event{ // happens every day
		StartDate: time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
		Frequency: Daily,
	}

	event2 := Event{ // happens every Sunday
		StartDate: time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
		Frequency: Weekly,
		WeekdayRule: WeekdayRule{
			Day:   time.Sunday,
			Occur: Every,
		},
	}

	event3 := Event{ // happens on the 15th of every month
		StartDate:  time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
		Frequency:  Monthly,
		DayOfMonth: 15,
	}

	event4 := Event{ // happens on the 15th of September every year
		StartDate:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2025, 10, 30, 0, 0, 0, 0, time.UTC),
		Frequency:  Yearly,
		Month:      9,
		DayOfMonth: 15,
	}

	event5 := Event{ // one-time event
		Date: time.Date(2023, 9, 15, 0, 0, 0, 0, time.UTC),
	}

	event6 := Event{ // happens on the first monday the month
		StartDate: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(2023, 7, 31, 0, 0, 0, 0, time.UTC),
		Frequency: Weekly,
		WeekdayRule: WeekdayRule{
			Day:   time.Monday,
			Occur: First,
		},
	}

	calendar := Calendar{
		Events: []Event{event1, event2, event3, event4, event5, event6},
	}

	filteredEvents := calendar.ListEventsBetween(start, end)

	if !slices.Contains(filteredEvents, event1) {
		t.Error("Event1 should be in the list")
	}

	if !slices.Contains(filteredEvents, event2) {
		t.Error("Event2 should be in the list")
	}

	if !slices.Contains(filteredEvents, event3) {
		t.Error("Event3 should be in the list")
	}

	if !slices.Contains(filteredEvents, event4) {
		t.Error("Event4 should be in the list")
	}

	if !slices.Contains(filteredEvents, event5) {
		t.Error("Event5 should be in the list")
	}

	if len(filteredEvents) != 5 {
		t.Errorf("There should be 3 events in the list, got %d", len(filteredEvents))
	}
}
