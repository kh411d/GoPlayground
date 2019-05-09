package main

import "fmt"

func maxDifference(a []int32) int32 {
	// Write your code here
	aa := a[1:]
	var s []int32
	var diff int32
	diff = -1
	for _, v := range aa {
		if len(s) > 0 {

			for _, w := range s {
				x := v - w
				if x > diff {
					diff = x
				}

			}
		}
		s = append(s, v)

	}
	return diff

}

func main() {
	a := []int32{5, 10, 8, 7, 6, 5}
	fmt.Print(maxDifference(a))

}
