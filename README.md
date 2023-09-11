# Go In-Memory Event Schedule

## Overview

The `schedule` Go module provides a simple yet flexible way to manage an in-memory event schedule. It allows you to schedule both one-off and recurrent events with ease. The core functionality is encapsulated in the `Event` struct, which can be embedded into any other structure you're trying to schedule.

**No Third-Party Dependencies**: This package is standalone and doesn't require any third-party libraries.

## Installation

```bash
go get github.com/edsonmedina/schedule
```

## Usage

### Import the Package

```go
import "github.com/edsonmedina/schedule"
```

### Create an Event

#### One-Time Event

```go
oneTimeEvent := schedule.Event{
    Date: time.Date(2023, 9, 5, 0, 0, 0, 0, time.UTC),
}
```

#### Recurring Daily Event

```go
dailyEvent := schedule.Event{
    Frequency: schedule.Daily,
    StartDate: time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
    EndDate:   time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
}
```

#### Recurring Weekly Event

```go
weeklyEvent := schedule.Event{
    Frequency: schedule.Weekly,
    WeekdayRule: schedule.WeekdayRule{
        Day:   time.Wednesday,
        Occur: schedule.Every,
    },
    StartDate: time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
    EndDate:   time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC),
}
```

### Create a Calendar and Add Events

```go
calendar := schedule.Calendar{
    Events: []schedule.Event{
        oneTimeEvent,
        dailyEvent,
        weeklyEvent,
        // ... other events
    },
}
```

### List Events Between Two Dates

```go
start := time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC)

events := calendar.ListEventsBetween(start, end)
```

## Features

- **Flexible Scheduling**: Supports daily, weekly, monthly, yearly, weekdays, and weekends events.
- **Weekday Rules**: Define complex rules for events that happen on specific weekdays.
- **Date Range**: Specify start and end dates for recurrent events.

## API Reference

### Event Struct

- `Date`: The date of the event. Use this field only for one-time events.
- `Frequency`: The frequency of the event (Daily, Weekly, Monthly, Yearly, Weekdays, Weekends).
- `DayOfMonth`: The day of the month the event occurs (e.g., 5 for the 5th day).
- `WeekdayRule`: Complex rules for weekday events (e.g., last Friday of each month).
- `Month`: The month the event occurs (e.g., 11 for November).
- `StartDate`: Start date for recurrent events. This field is mandatory for recurring events.
- `EndDate`: End date for recurrent events. This field is mandatory for recurring events.

### Calendar Struct

- `Events`: A slice of `Event` structs.

### Methods

- `ListEventsBetween(start, end time.Time) []Event`: Lists events that happen between the given start and end dates.

## Contributing

Feel free to open issues or submit pull requests.
