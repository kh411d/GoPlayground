package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	i, err := strconv.ParseInt("1686633808139", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i/1000, 0).UTC().Format(time.RFC3339)
	fmt.Println(tm)
}

//1590683330545
