package events

import (
	"fmt"
	"time"
)

type CalendarColumn struct {
	Date   time.Time
	Name   string
	Events []Event
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
			Name: fmt.Sprint("Day ", i),
			Events: []Event{
				{Task: &Task{Title: "skibidi"}, From: time.Now(), To: time.Now().Add(time.Hour)},
				{Task: &Task{Title: "bip bop"}, From: time.Now(), To: time.Now().Add(time.Hour)},
			},
		}
		CalendarColumns = append(CalendarColumns, newColumn)
	}
}
