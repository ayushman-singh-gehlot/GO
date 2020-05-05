package main

import (
	"fmt"
	"time"
)
var status string
func main() {
	timeHour:=time.Now().Hour()
	timeMinute:=time.Now().Minute()
	println("current time :- ",timeHour,":",timeMinute)
	if timeHour>=0 && timeHour<12{
		status="good morning"
	}else if timeHour>=12 && timeHour<17{
		status="good afternoon"
	}else if timeHour>=17 && timeHour<21{
		status="good evening"
	}else if timeHour>=21 && timeHour<=23{
		status="good night"
	}
	fmt.Println(status)
}
