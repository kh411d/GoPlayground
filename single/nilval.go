package main

import "fmt"

type xxx struct {
	id *int
}

func main() {

	x := xxx{}

	fmt.Println(len(x.id))
}
