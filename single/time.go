package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//test doang lagi
	// then := now.Add(-179 * time.Second)
	// fmt.Println(now)
	// fmt.Println(then)

	dtstr := "2021-04-27T11:19:36"
	layout := "2006-01-02T15:04:05"

	loc, _ := time.LoadLocation("Asia/Riyadh")
	t := now.In(loc)
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", time.Now().UTC())

	//timeUTC := time.Now().UTC()
	timeInFuture, _ := time.Parse(layout, dtstr)
	fmt.Printf("%v\n", timeInFuture)

	starttime := time.Now()
	flagged := starttime.Add(-1 * time.Duration(5) * time.Minute)
	fmt.Printf("%v\n", starttime)
	fmt.Printf("%v\n", flagged)

	zeroTime := time.Time{}
	fmt.Printf("%v\n", zeroTime.IsZero())

	//xtime := "2021-10-03T01:15:00+07:00"
	xtime := "2021-09-30T12:40:00+07:00"
	xtimeInFuture, _ := time.Parse(time.RFC3339, xtime)
	fmt.Println(xtimeInFuture.UTC())
	fmt.Println(xtimeInFuture.In(loc))
	locJ, _ := time.LoadLocation("Asia/Jakarta")
	fmt.Println(xtimeInFuture.In(locJ))

	ytime := "2021-10-03T01:30:00+07:00"
	ytimeInFuture, _ := time.Parse(time.RFC3339, ytime)
	fmt.Println(ytimeInFuture.In(loc))

	c, _ := time.Parse("2006-01-02 15:04:05", "2021-02-25 12:32:01")
	fmt.Println(c.Format(time.RFC3339))

}
