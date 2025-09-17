package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	loc, _ := time.LoadLocation("Asia/Dubai")
	fmt.Println(t.In(loc))
	fmt.Println(t.In(loc).AddDate(0, 0, 1))

	tp, _ := time.ParseInLocation("2006-01-02", "2024-06-30", loc)
	fmt.Println(tp.AddDate(0, 0, 1))

	locf := time.FixedZone("GMT", 4)
	fmt.Println(t.In(locf))

}
