package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	start := time.Now()
	time.Sleep(time.Millisecond * 500)
	stop := time.Now()

	latency := time.Since(start)
	l := stop.Sub(start)
	fmt.Printf("%#v\n", int64(l))
	fmt.Println(strconv.FormatInt(int64(l), 10))
	fmt.Println(strconv.FormatInt(int64(latency), 10))
	fmt.Println(l)
	fmt.Println(latency)
}
