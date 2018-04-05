package main

import (
	"fmt"
	"time"
)

func testaja() {
	defer count()
	time.Sleep(5 * time.Second)
	end := time.Now()

	fmt.Printf("end %v\n", end)
}

func count() {
	start := time.Now()
	fmt.Printf("start %v\n", start)
}

func main() {
	testaja()
}
