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
	Id       string    `json:"_id" bson:"_id"` // This is important so the attribute `Id` is renamed to _id in mongo (mongo renames using bson)
}

type CalendarEntry struct {
    Start string   		`json:"start"`	
    End string			`json:"end"`	
    EventName string 	`json:"eventname"`	
    Notes string		`json:"notes"`	
    Pass string			`json:"pass"`	
    Color string 		`json:"color"`	
}