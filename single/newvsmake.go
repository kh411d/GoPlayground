package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing ground local")

	type A struct {
		a string
		b string
	}

	w := make(A)
	fmt.Println(w)

	//Equivalent
	x := new(A)
	x.a = "ini a"
	fmt.Println(x)

	//Equivalent
	y := &A{a: "ini a"}
	fmt.Println(y)

	//Equivalent
	var z *A
	z = &A{}
	fmt.Println(z)

	var p *[]int = new([]int)    // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 1) // the slice v now refers to a new array of 100 ints

	fmt.Println(*p)
	fmt.Println(v)

}
