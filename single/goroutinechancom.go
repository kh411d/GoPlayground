package main

import "fmt"

//https://golang.org/ref/mem

/*var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	//c <- 0
	close(c)
}

func main() {
	go f()
	x := <-c
	print(a)
	println(x)
	fmt.Scanln("")
}*/

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	<-c
}
func main() {
	go f()
	c <- 0
	println(a)
	fmt.Scanln("")
}
