package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// loopDate := now.AddDate(0, 0, 1)
	// t1, _ := time.Parse(time.RFC3339, loopDate.Format(`2006-01-02T`)+`15:00:00`+"+07:00")
	// t2, _ := time.Parse(time.RFC3339, loopDate.Format(`2006-01-02T`)+`17:00:00`+"+07:00")
	// fmt.Println(now)
	// fmt.Println(loopDate)
	// fmt.Println(t1)
	// fmt.Println(t2)
	// if now.After(t1) && now.Before(t2) { // In the delivery slot
	//     fmt.Println("asdf")
	// } else if now.After(t2) { // After the delivery slot
	//     fmt.Println("qwer")
	// }

	//xtime := "2021-10-03T01:15:00+07:00"
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// }, {
	//     "start_time": "12/23/2023 9:00:00 AM",
	//     "end_time": "12/23/2023 9:30:00 AM"
	//   }, {
	//     "start_time": "12/23/2023 9:30:00 AM",
	//     "end_time": "12/23/2023 10:00:00 AM"
	//   }, {
	//     "start_time": "12/23/2023 10:00:00 AM",
	//     "end_time": "12/23/2023 10:30:00 AM"
	//   },

	xtime := "2024-08-17T10:00:00+04:00"
	xtimeInFuture, _ := time.Parse(time.RFC3339, xtime)
	fmt.Println(xtimeInFuture.UTC().Format(time.RFC3339))
	fmt.Println(xtimeInFuture.In(loc).Format(time.RFC3339))
	fmt.Println("")

	// ytime := time.Now().UTC().Format(time.RFC3339)
	// ytimeInFuture, _ := time.Parse(time.RFC3339, ytime)
	// fmt.Println(ytimeInFuture.UTC().Format(time.RFC3339))
	// fmt.Println(ytimeInFuture.In(loc).Format(time.RFC3339))

	//2024-02-22T07:15:00Z
	//2024-02-22T14:15:00+07:00

}

// 2023-11-27T20:00:00Z
// 2023-11-25T19:55:00Z

//2023-12-22T08:00:00.000Z
