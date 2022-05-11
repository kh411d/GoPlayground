package main

import (
	"fmt"
)

func main() {
	x := make(map[string]int64)
	x["name"] = 2
	k := x["baba"]
	fmt.Println(k)

	fmt.Printf("ini v %#v", x["baba"] == int64(3))
}
