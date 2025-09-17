package main

import (
	"fmt"
	"time"
)

func main() {
	quickAdvanceDay := 31
	st := time.Now()

	et := st.AddDate(0, 0, quickAdvanceDay+1).Add(-1 * time.Nanosecond)

	fmt.Println(st.Format(time.RFC3339))
	fmt.Println(et.Format(time.RFC3339))

	previousDate := st.AddDate(0, 0, -(quickAdvanceDay + 1))
	nextDate := st.AddDate(0, 0, quickAdvanceDay+1)
	fmt.Printf("Previous Date : %s\n", previousDate.Format(time.RFC3339))
	fmt.Printf("End Date: %s\n", nextDate.Format(time.RFC3339))

	diff := et.Sub(st)

	x := int((diff.Hours() / 24))
	fmt.Println(x)

}
