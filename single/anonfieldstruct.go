package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

type Human struct {
	*Animal
	Address string
	ID      string
}

func main() {
	a := new(Animal)
	x := Human{a, "", ""}

	x.Name = "kambing"
	fmt.Println(x.Name)
}
