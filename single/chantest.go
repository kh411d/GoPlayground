package main

import (
	"fmt"
	"time"
)

func main() {
	x := make(chan string)
	// output := func() {
	// 	go func() {
	// 		<-x
	// 		fmt.Printf("its done")

	// 	}()
	// }
	// output()

	go func(x chan string) {
		fmt.Printf("its done")
		fmt.Println(<-x)

	}(x)

	go func(x chan string) {
		fmt.Printf("its done2")
		fmt.Println(<-x)

	}(x)

	time.Sleep(5 * time.Second)
	x <- "done"
}
