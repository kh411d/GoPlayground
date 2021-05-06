package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Hello, playground")
	timeNow := time.Now()
	randig := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(201)
	fmt.Println(timeNow.Unix())
	fmt.Println(randig)
	fmt.Printf("%d%d\n", time.Now().Unix(), randig)
}

