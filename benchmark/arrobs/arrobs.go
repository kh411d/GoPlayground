package main

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
)

type observee []int

func (o *observee) observer(fn func(observee), td time.Duration) chan<- struct{} {
	done := make(chan struct{})
	go func() {
		var pastLen int
		for {
			select {
			case <-done:
				return
			case <-time.NewTicker(td).C:
				nowLen := len(*o)
				if nowLen != pastLen {
					fn(*o)
				}
				pastLen = nowLen
			}
		}
	}()
	return done
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	defer profile.Start(profile.ProfilePath("/Users/khalidity/go/src/github.com/kh411d/MyGoPlayground/benchmark/arrobs/")).Stop()
	var theArray observee
	var total int
	t := theArray.observer(func(o observee) {
		//fmt.Printf("Observe %v\n", o)
		total = sum(o)
	}, 1*time.Nanosecond)

	for i := 1; i <= 100000; i++ {
		theArray = append(theArray, i)
		time.Sleep(5 * time.Nanosecond)
	}
	t <- struct{}{}
	fmt.Println("Ticker/Observer stopped")
	fmt.Printf("Total count: %v\n", total)
	//fmt.Printf("The final array: %v\n", theArray)
}
