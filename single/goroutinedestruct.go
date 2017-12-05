package main

import "fmt"

var limit = make(chan int, 2)

func main() {
	work := []func(){
		func() { println("1") },
		func() { println("2") },
		func() { println("3") },
	}
	for _, w := range work {
		go func(w func()) {
			limit <- 1
			w()
			<-limit
		}(w)
	}
	//select {}
	fmt.Scanln("")
}
