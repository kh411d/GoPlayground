package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing ground local")
	//x := []int{}
	x := make([]int,5)

	fmt.Println(x)

	y := [3]string{"Лайка", "Белка", "Стрелка"}
	fmt.Println(y[:]) 
}
