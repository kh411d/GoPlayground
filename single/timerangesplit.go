package main

import (
	"fmt"
	"time"
)

type SplitTime struct {
	from time.Time
	to   time.Time
}

func SplitByDays(fromTime, toTime time.Time, nDays int) (timeBucket []SplitTime) {
	days := int(toTime.Sub(fromTime).Hours()) / 24
	nMonth := days / nDays
	nDay := days % nDays

	for i := 0; i < nMonth; i++ {
		if l := len(timeBucket); l > 0 {
			timeBucket = append(timeBucket, SplitTime{
				from: timeBucket[l-1].to.Add(5 * time.Minute), //add 5 minute to prevent duplicate
				to:   timeBucket[l-1].to.Add(30 * 24 * time.Hour),
			})
		} else {
			timeBucket = append(timeBucket, SplitTime{
				from: fromTime,
				to:   fromTime.Add(30 * 24 * time.Hour),
			})
		}
	}

	//get the leftover days
	if nDay > 0 {
		if l := len(timeBucket); l > 0 {
			timeBucket = append(timeBucket, SplitTime{
				from: timeBucket[l-1].to.Add(5 * time.Minute), //add 5 minute to prevent duplicate
				to:   toTime,
			})
		} else {
			timeBucket = append(timeBucket, SplitTime{
				from: fromTime,
				to:   toTime,
			})
		}
	}

	return

}

func main() {
	x := time.Now()
	y := x.Add(32 * 24 * time.Hour)
	fmt.Println(x)
	fmt.Println(y)

	timeBucket := SplitByDays(x, y, 30)

	for _, v := range timeBucket {
		fmt.Printf("From: %s\n", v.from.Format(time.RFC3339))
		fmt.Printf("To: %s\n", v.to.Format(time.RFC3339))
		fmt.Println("--------")
	}

}

// func main() {
// 	x := time.Now()
// 	y := x.Add(32 * 24 * time.Hour)
// 	fmt.Println(x)
// 	fmt.Println(y)

// 	days := int(y.Sub(x).Hours()) / 24
// 	nMonth := days / 30
// 	nDay := days % 30

// 	var timeBucket []SplitTime
// 	for i := 0; i < nMonth; i++ {
// 		if l := len(timeBucket); l > 0 {
// 			timeBucket = append(timeBucket, SplitTime{
// 				from: timeBucket[l-1].to.Add(1 * 24 * time.Hour),
// 				to:   timeBucket[l-1].to.Add(30 * 24 * time.Hour),
// 			})
// 		} else {

// 			timeBucket = append(timeBucket, SplitTime{
// 				from: x,
// 				to:   x.Add(30 * 24 * time.Hour),
// 			})
// 		}
// 	}

// 	if nDay > 0 {
// 		if l := len(timeBucket); l > 0 {
// 			timeBucket = append(timeBucket, SplitTime{
// 				from: timeBucket[l-1].to.Add(1 * 24 * time.Hour),
// 				to:   y,
// 			})
// 		} else {

// 			timeBucket = append(timeBucket, SplitTime{
// 				from: x,
// 				to:   y,
// 			})
// 		}
// 	}

// 	fmt.Println(nMonth)
// 	fmt.Println(nDay)

// 	for _, v := range timeBucket {
// 		fmt.Printf("From: %s\n", v.from.Format(time.RFC3339))
// 		fmt.Printf("To: %s\n", v.to.Format(time.RFC3339))
// 		fmt.Println("--------")
// 	}

// }
