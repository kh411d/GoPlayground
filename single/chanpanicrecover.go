package main

import (
	"errors"
	"fmt"
	"time"
)

func HoldTimeout(touchTimeAfter time.Duration) func() {

	// prepare stopper
	stop := make(chan bool, 1)

	// Looper for chan
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(touchTimeAfter)
				fmt.Println("sleep")
			}
		}
	}()

	return func() {
		stop <- true
		fmt.Println("release")
	}
}

var testIds map[string]int = make(map[string]int)

func hold() {
	release := HoldTimeout(3 * time.Second)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		release()
	}()

	time.Sleep(5 * time.Second)
	if _, ok := testIds["panic"]; !ok {
		testIds["panic"] = 1
		panic(errors.New("test panic"))
	}
	return
}

func main() {
	hold()

	var input string
	fmt.Scanln(&input)
}
