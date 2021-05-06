package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup
	c1 := make(chan string)
	c2 := make(chan string)

	w.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}(&w)
	w.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}(&w)

	//for i := 0; i < 2; i++ {
	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	}
	//}
	w.Wait()
}
