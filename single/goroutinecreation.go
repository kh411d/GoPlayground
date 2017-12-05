package main

import "fmt"

var a string

/*func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	println("oo")
	go f()
}
*/
func hello() {
	go func() { a = "hello"; print(a) }()
	print(a)
}

func main() {
	a = "kambing"
	hello()
	var input string
	fmt.Scanln(&input)
}
