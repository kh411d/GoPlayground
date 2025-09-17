package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("stopped")
				return
			}
		}
	}()
	time.Sleep(3 * time.Second)
	stop <- true

	time.Sleep(1 * time.Second)
	fmt.Println("fill")
	stop <- true

	fmt.Println(stop)
	time.Sleep(1 * time.Second)
}
