package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}

	return duplicate_frequency
}

func main() {
	//timeNow := time.Now()
	var x []string
	for i := 1; i <= 999; i++ {
		time.Sleep(100 * time.Millisecond)
		randInt := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(201)
		//x = append(x, randInt)

		x = append(x, fmt.Sprintf("%d%03d", time.Now().Unix(), randInt))
	}
	//x := fmt.Sprintf("%s %03d", timeNow.Format("150405"), randInt)

	//fmt.Println(x)

	dup_map := dup_count(x)

	//fmt.Println(dup_map)

	for k, v := range dup_map {
		if v > 1 {
			fmt.Printf("Item : %s , Count : %d\n", k, v)
		}
	}
}
