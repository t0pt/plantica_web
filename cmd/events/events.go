package events

import (
	"time"
)

type CalendarColumn struct {
	Date   time.Time
	Events map[int][]Event
	Tasks  []Task
}

type Task struct { // tasks are saved for the whole day
	Date  time.Time // used for keeping the date only
	Title string
	Color string
}

type Event struct {
	*Task
	From time.Time
	To   time.Time
}

var CalendarColumns = []CalendarColumn{}

func CreateCalendarColumns() {
	CalendarColumns = []CalendarColumn{}
	for i := range []int{0, 1, 2, 3, 4} {
		newColumn := CalendarColumn{
			Date: time.Now().AddDate(0, 0, i),
			Events: map[int][]Event{
				18: {{Task: &Task{Title: "skibidi", Color: "green"}, From: time.Now(), To: time.Now().Add(time.Hour)},
					{Task: &Task{Title: "skibidi", Color: "orange"}, From: time.Now(), To: time.Now().Add(time.Hour)}},
				20: {{Task: &Task{Title: "bip bop"}, From: time.Now(), To: time.Now().Add(time.Hour)}},
			},
			Tasks: []Task{
				{Title: "foo", Color: "red"},
				{Title: "baar", Color: "blue"},
			},
		}
		CalendarColumns = append(CalendarColumns, newColumn)
	}
}
