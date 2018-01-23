package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {

			fmt.Println("processed:", m)
			time.Sleep(time.Second * 3)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed
}
