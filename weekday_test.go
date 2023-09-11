package schedule

import (
	"testing"
	"time"
)

func TestWeekdayRule_HappensOn(t *testing.T) {
	tests := []struct {
		name string
		rule WeekdayRule
		date time.Time
		want bool
	}{
		{
			name: "Every Monday",
			rule: WeekdayRule{Day: time.Monday, Occur: Every},
			date: time.Date(2023, 9, 11, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "Last Monday of the month - true",
			rule: WeekdayRule{Day: time.Monday, Occur: Last},
			date: time.Date(2022, 11, 28, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "Last Monday of the month - false",
			rule: WeekdayRule{Day: time.Monday, Occur: Last},
			date: time.Date(2022, 11, 21, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "First Monday of the month - true",
			rule: WeekdayRule{Day: time.Monday, Occur: First},
			date: time.Date(2022, 11, 7, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "First Monday of the month - false",
			rule: WeekdayRule{Day: time.Monday, Occur: First},
			date: time.Date(2022, 11, 8, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "Second Monday of the month- true",
			rule: WeekdayRule{Day: time.Monday, Occur: Second},
			date: time.Date(2022, 11, 14, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "Second Monday of the month - false",
			rule: WeekdayRule{Day: time.Monday, Occur: Second},
			date: time.Date(2022, 11, 15, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "Third Monday of the month - true",
			rule: WeekdayRule{Day: time.Monday, Occur: Third},
			date: time.Date(2022, 11, 21, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "Third Monday of the month - false",
			rule: WeekdayRule{Day: time.Monday, Occur: Third},
			date: time.Date(2022, 11, 22, 0, 0, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rule.HappensOn(tt.date); got != tt.want {
				t.Errorf("WeekdayRule.HappensOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
