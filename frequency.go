package schedule

type Frequency string

const (
	Daily    Frequency = "daily"
	Weekly   Frequency = "weekly"
	Monthly  Frequency = "monthly"
	Yearly   Frequency = "yearly"
	Weekdays Frequency = "weekdays"
	Weekends Frequency = "weekends"
)
