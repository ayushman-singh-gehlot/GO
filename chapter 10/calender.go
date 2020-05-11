package main

import (
	"fmt"
	"log"
	"packages/calender"
)

func addEvent(title string, year, month, day int) *calender.Event {
	event := calender.Event{}
	err := event.SetTitle(title)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(year)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(month)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(day)
	if err != nil {
		log.Fatal(err)
	}
	return &event

}

func displayEvent(event *calender.Event) {
	fmt.Println(event.Title())
	fmt.Println(event.Day(), event.Month(), event.Year())

}

func main() {
	event1 := addEvent("My birthday :)", 2019, 10, 18) // returning address of the struct pointer Event
	displayEvent(event1)
}
