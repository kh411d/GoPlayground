package main

import (
	"fmt"
	"sync"
)

var a string
var once sync.Once

func setup() {
	a = "setup ok"
}

func doprint() {
	once.Do(setup)
	println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	fmt.Scanln()
}
