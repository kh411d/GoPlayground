package main

import (
	"fmt"
	"time"
)

func main() {

	timeNow := time.Now()
	fmt.Printf("Current Time: %v\n", timeNow.Format(time.RFC3339))

	appTime := timeNow.Add((time.Hour * 2))
	reminderTimeLimit := appTime.Add(time.Minute * -10)

	fmt.Printf("Appointment time: %v\n", appTime.Format(time.RFC3339))
	fmt.Printf("Appointment time limit: %v\n", reminderTimeLimit.Format(time.RFC3339))
	reminderTime := appTime.Add(time.Minute * -30)

	fmt.Printf("Reminder time: %v\n", reminderTime.Format(time.RFC3339))
	if reminderTime.Before(reminderTimeLimit) && reminderTime.After(timeNow) {
		fmt.Println("ok to remind")
	}
}
