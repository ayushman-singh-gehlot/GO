package calender

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(titleName string) error {
	if utf8.RuneCountInString(titleName) > 30 {
		return errors.New("invalid title")
	}
	e.title = titleName
	return nil
}

type Date struct { // this is struct storing date
	year  int // all the fields are unexported
	month int
	day   int
}

// getter methods
// Year returns value of year
func (d *Date) Year() int {
	return d.year
}
func (d *Date) Month() int {
	return d.month
}
func (d *Date) Day() int {
	return d.day
}

// setter methods
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid Year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
