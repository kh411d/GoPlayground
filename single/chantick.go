package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	done := make(chan string)
	go func(d time.Duration, f func(time.Time)) {
		// for x := range time.Tick(d) {
		// 	f(x)
		// }
		for {
			select {
			case <-done:
				f(time.Now())
			case x := <-time.Tick(d):
				f(x)
			}
		}

	}(1*time.Second, func(time.Time) {
		i++
		fmt.Println(i)
	})

	time.Sleep(5 * time.Second)
	done <- "done"
}
