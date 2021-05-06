package main

import (
	"fmt"
	"time"
)

func main() {
	x := []string{"01", "02", "03", "04", "05", "06", "07"}

	i := 0
	for {
		fmt.Println(x[i])
		i++
		if i > len(x)-1 {
			i = 0
		}
		time.Sleep(1 * time.Second)
	}

}
