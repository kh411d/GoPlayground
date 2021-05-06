package main

import (
	"fmt"
	"time"
)

var timeNow time.Time

func main() {
	timeNow = time.Now()
	a := makeTimestamp()

	fmt.Printf("%d \n", a)
	fmt.Printf("%d\n", timeNow.Unix())
}

/*
func makeTimestamp() int64 {
    return timeNow.UnixNano() / int64(time.Millisecond)
}*/

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
