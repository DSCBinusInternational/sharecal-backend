package models

import "time"

type TimeEntry struct {
	Name  string      `json:"name"`
	Time  []time.Time `json:"time"`
	Notes string      `json:"notes"`
	Color string      `json:"color"`
}

type DayEntry map[int][]TimeEntry
type MonthEntry map[int]DayEntry
type YearEntry map[int]MonthEntry

type Calendar struct {
	Data     YearEntry `json:"data"`
	Passcode string    `json:"passcode"`
	Id       string    `json:"_id"`
}
