package main

import (
	"fmt"
	"sort"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var temp int
	temp = 1
	sort.Ints(A)
	for _, v := range A {

		if v == temp {
			temp++
		}
		/*if v == last {
			continue
		}

		if v != last+1 {
			if last+1 > temp {
				temp = last + 1
			}
		} else if temp == 0 {
			temp = last + 1
		}*/

		//last = v

	}
	return temp
}

func main() {
	x := []int{1, 3, 6, 4, 1, 2}
	//x := []int{1, 2, 3}
	fmt.Println(Solution(x))
}
