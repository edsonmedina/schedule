package schedule

import "time"

type Calendar struct {
	Events []Event
}

func (c *Calendar) ListEventsBetween(start, end time.Time) []Event {
	var filteredEvents []Event

	for _, event := range c.Events {
		current := start
		for current.Before(end) || current.Equal(end) {
			if event.HappensOn(current) {
				filteredEvents = append(filteredEvents, event)
				break // No need to check further dates for this event
			}
			current = current.AddDate(0, 0, 1) // Increment by one day
		}
	}

	return filteredEvents
}
