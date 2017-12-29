package main

import (
	"fmt"
)

type HandlerFunc func(string, int)

func (f HandlerFunc) ServeHTTP(w string, r int) {
	f(w, r)
}

func HFunc(handler func(string, int)) {
	x := HandlerFunc(handler)
	x.ServeHTTP("kambing", 1)
}

func Joni(i string, x int) {
	fmt.Println(i)
	fmt.Println(x)
}

func main() {
	//	HFunc(Joni)

	HandlerFunc(Joni).ServeHTTP("kambing", 1)

}
