package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//test doang
	// then := now.Add(-179 * time.Second)
	// fmt.Println(now)
	// fmt.Println(then)

	dtstr := "2021-04-27T11:19:36"
	layout := "2006-01-02T15:04:05"

	loc, _ := time.LoadLocation("Asia/Jakarta")
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
	fmt.Printf("%v", zeroTime.IsZero())

}
