package models

import "time"

type TimeEntry struct {
	Name  string      `json:"name"`
	Time  []time.Time `json:"time"`
	Notes string      `json:"notes"`
	Color string      `json:"color"`
}

type DayEntry map[string][]TimeEntry
type MonthEntry map[string]DayEntry
type YearEntry map[string]MonthEntry

type Calendar struct {
	Data     YearEntry `json:"data"`
	Passcode string    `json:"passcode"`
}
