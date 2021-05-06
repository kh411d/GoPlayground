package main

import "fmt"

func fn(m map[int]int, s *[]int) {
	//m = make(map[int]int)
	m[1] = 1
	*s = append(*s, 2)
	fmt.Printf("map inside fn %v\n", m)
	fmt.Printf("slice inside fn %v\n", s)
}

func main() {
	var m map[int]int
	var s []int
	m = make(map[int]int)
	s = []int{1}
	fn(m, &s)
	fmt.Printf("map outside %v\n", m)
	fmt.Printf("slice outside %v\n", s)
}

/*
//https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
func fn(m map[int]int) {
        m = make(map[int]int)
}

func main() {
        var m map[int]int
        fn(m)
        fmt.Println(m == nil)
}
*/
