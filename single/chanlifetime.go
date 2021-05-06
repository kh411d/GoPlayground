package main

import (
	"fmt"
)

func main() {
	x := make(chan bool)
	y := make(chan bool)
	//var w sync.WaitGroup

	//w.Add(1)
	//channel reader
	//select will make this go routine wait for a channel action case
	go func(x chan bool, y chan bool) {
		//defer wg.Done()
		fmt.Println("READ xy")
		select {
		case <-x:
			fmt.Println("RECEIVE x")
		case <-y:
			fmt.Println("RECEIVE y")
		}
	}(x, y)

	//this channel will race the y chanel to be read by the reader
	//though this routine will die when main thread ended
	//so deadlock will not happening as the main thread ended
	go func() {
		fmt.Println("SENT x")
		x <- true
	}()

	//this channel will race the x chanel to be read by the reader
	//This will block until the reader receive the channel
	//if the other channel x got read by the reader first then
	//deadlock will happen, because the reader is already ended or gone
	fmt.Println("SENT y")
	y <- true
}

// go func() {
// 	y <- true
// }()

//w.Wait()

//panic("show stack")
